package conf

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	kc "github.com/infraboard/keyauth/client"
	"github.com/infraboard/mcube/cache/memory"
	"github.com/infraboard/mcube/cache/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db        *sql.DB
	mgoclient *mongo.Client
)

func newConfig() *Config {
	return &Config{
		App:     newDefaultAPP(),
		HTTP:    newDefaultHTTP(),
		GRPC:    newDefaultGRPC(),
		Log:     newDefaultLog(),
		MySQL:   newDefaultMySQL(),
		Mongo:   newDefaultMongoDB(),
		Cache:   newDefaultCache(),
		Keyauth: newDefaultKeyauth(),
	}
}

// Config 应用配置
type Config struct {
	App     *app     `toml:"app"`
	HTTP    *http    `toml:"http"`
	GRPC    *grpc    `toml:"grpc"`
	Log     *log     `toml:"log"`
	MySQL   *mysql   `toml:"mysql"`
	Mongo   *mongodb `toml:"mongodb"`
	Keyauth *keyauth `toml:"keyauth"`
	Cache   *_cache  `toml:"cache"`
}

type app struct {
	Name string `toml:"name" env:"APP_NAME"`
	Key  string `toml:"key" env:"APP_KEY"`
}

func newDefaultAPP() *app {
	return &app{
		Name: "demo",
		Key:  "default",
	}
}

type http struct {
	Host      string `toml:"host" env:"HTTP_HOST"`
	Port      string `toml:"port" env:"HTTP_PORT"`
	EnableSSL bool   `toml:"enable_ssl" env:"HTTP_ENABLE_SSL"`
	CertFile  string `toml:"cert_file" env:"HTTP_CERT_FILE"`
	KeyFile   string `toml:"key_file" env:"HTTP_KEY_FILE"`
}

func (a *http) Addr() string {
	return a.Host + ":" + a.Port
}

func newDefaultHTTP() *http {
	return &http{
		Host: "127.0.0.1",
		Port: "8050",
	}
}

type grpc struct {
	Host      string `toml:"host" env:"GRPC_HOST"`
	Port      string `toml:"port" env:"GRPC_PORT"`
	EnableSSL bool   `toml:"enable_ssl" env:"GRPC_ENABLE_SSL"`
	CertFile  string `toml:"cert_file" env:"GRPC_CERT_FILE"`
	KeyFile   string `toml:"key_file" env:"GRPC_KEY_FILE"`
}

func (a *grpc) Addr() string {
	return a.Host + ":" + a.Port
}

func newDefaultGRPC() *grpc {
	return &grpc{
		Host: "127.0.0.1",
		Port: "18050",
	}
}

type log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
}

func newDefaultLog() *log {
	return &log{
		Level:   "debug",
		PathDir: "logs",
		Format:  "text",
		To:      "stdout",
	}
}

// Auth auth 配置
type keyauth struct {
	Host         string `toml:"host" env:"KEYAUTH_HOST"`
	Port         string `toml:"port" env:"KEYAUTH_PORT"`
	ClientID     string `toml:"client_id" env:"KEYAUTH_CLIENT_ID"`
	ClientSecret string `toml:"client_secret" env:"KEYAUTH_CLIENT_SECRET"`
}

func (a *keyauth) Addr() string {
	return a.Host + ":" + a.Port
}

func (a *keyauth) Client() (*kc.Client, error) {
	if kc.C() == nil {
		conf := kc.NewDefaultConfig()
		conf.SetAddress(a.Addr())
		conf.SetClientCredentials(a.ClientID, a.ClientSecret)
		client, err := kc.NewClient(conf)
		if err != nil {
			return nil, err
		}
		kc.SetGlobal(client)
	}

	return kc.C(), nil
}

func newDefaultKeyauth() *keyauth {
	return &keyauth{}
}

func newDefaultMongoDB() *mongodb {
	return &mongodb{
		Database:  "",
		Endpoints: []string{"127.0.0.1:27017"},
	}
}

type mongodb struct {
	Endpoints []string `toml:"endpoints" env:"MONGO_ENDPOINTS" envSeparator:","`
	UserName  string   `toml:"username" env:"MONGO_USERNAME"`
	Password  string   `toml:"password" env:"MONGO_PASSWORD"`
	Database  string   `toml:"database" env:"MONGO_DATABASE"`
}

// Client 获取一个全局的mongodb客户端连接
func (m *mongodb) Client() *mongo.Client {
	if mgoclient == nil {
		panic("please load mongo client first")
	}

	return mgoclient
}

func (m *mongodb) GetDB() *mongo.Database {
	return m.Client().Database(m.Database)
}

func (m *mongodb) getClient() (*mongo.Client, error) {
	opts := options.Client()

	cred := options.Credential{
		AuthSource: m.Database,
	}

	if m.UserName != "" && m.Password != "" {
		cred.Username = m.UserName
		cred.Password = m.Password
		cred.PasswordSet = true
		opts.SetAuth(cred)
	}
	opts.SetHosts(m.Endpoints)
	opts.SetConnectTimeout(5 * time.Second)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, fmt.Errorf("new mongodb client error, %s", err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("ping mongodb server(%s) error, %s", m.Endpoints, err)
	}

	return client, nil
}

type mysql struct {
	Host        string `toml:"host" env:"MYSQL_HOST"`
	Port        string `toml:"port" env:"MYSQL_PORT"`
	UserName    string `toml:"username" env:"MYSQL_USERNAME"`
	Password    string `toml:"password" env:"MYSQL_PASSWORD"`
	Database    string `toml:"database" env:"MYSQL_DATABASE"`
	MaxOpenConn int    `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int    `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int    `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	lock        sync.Mutex
}

func newDefaultMySQL() *mysql {
	return &mysql{
		Database:    "demo",
		Host:        "127.0.0.1",
		Port:        "3306",
		MaxOpenConn: 200,
		MaxIdleConn: 16,
		MaxLifeTime: 300,
	}
}

// getDBConn use to get db connection pool
func (m *mysql) getDBConn() (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true", m.UserName, m.Password, m.Host, m.Port, m.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", dsn, err.Error())
	}
	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	db.SetConnMaxLifetime(time.Minute * time.Duration(m.MaxLifeTime))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}
	return db, nil
}
func (m *mysql) GetDB() (*sql.DB, error) {
	// 加载全局数据量单例
	m.lock.Lock()
	defer m.lock.Unlock()
	if db == nil {
		conn, err := m.getDBConn()
		if err != nil {
			return nil, err
		}
		db = conn
	}
	return db, nil
}

func newDefaultCache() *_cache {
	return &_cache{
		Type:   "memory",
		Memory: memory.NewDefaultConfig(),
		Redis:  redis.NewDefaultConfig(),
	}
}

type _cache struct {
	Type   string         `toml:"type" json:"type" yaml:"type" env:"CACHE_TYPE"`
	Memory *memory.Config `toml:"memory" json:"memory" yaml:"memory"`
	Redis  *redis.Config  `toml:"redis" json:"redis" yaml:"redis"`
}
