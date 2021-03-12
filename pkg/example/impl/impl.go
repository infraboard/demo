package impl

import (
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
}

func (s *service) Config() error {
	// get global config with here

	return nil
}

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return example.HttpEntry()
}

func init() {
	pkg.RegistryService("example", Service)
}
