package example

// NewBook todo
func NewBook(req *CreateBookRequest) *Book {
	return &Book{
		Id:   "mock id",
		Name: req.Name,
	}
}

// NewBookSet 实例
func NewBookSet() *BookSet {
	return &BookSet{
		Items: []*Book{},
	}
}
