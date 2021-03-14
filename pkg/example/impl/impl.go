package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pb/http"

	"github.com/infraboard/demo/pkg"
	"github.com/infraboard/demo/pkg/example"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	example.UnimplementedServiceServer

	log logger.Logger
}

func (s *service) Config() error {
	// get global config with here
	s.log = zap.L().Named("Example")
	return nil
}

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return example.HttpEntry()
}

func init() {
	pkg.RegistryService("example", Service)
}
