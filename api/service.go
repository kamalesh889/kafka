package api

type Service interface {
	CreateProduct(*ProductRequest) error
	CreateUser(*UserRequest) error
}

func (s *server) CreateProduct(productDetails *ProductRequest) error {
	return nil
}

func (s *server) CreateUser(userDetails *UserRequest) error {
	return nil
}
