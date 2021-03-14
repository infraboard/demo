package example

import "github.com/infraboard/mcube/http/request"

// NewQueryBookRequest 查询book列表
func NewQueryBookRequest(page *request.PageRequest) *QueryBookRequest {
	return &QueryBookRequest{
		Page: &page.PageRequest,
	}
}
