package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/demo/client"
	"github.com/infraboard/demo/pkg"
	"github.com/infraboard/demo/pkg/example"
)

var (
	api = &handler{}
)

type handler struct {
	service example.ServiceClient
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("examples")

	r.BasePath("books")
	r.Handle("POST", "/", h.CreateBook)
	r.Handle("GET", "/", h.QueryBook)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.Example()
	return nil
}

func init() {
	pkg.RegistryHTTPV1("example", api)
}
