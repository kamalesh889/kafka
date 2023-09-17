package api

type UserRequest struct {
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}

type UserResponse struct {
	Name string `json:"name"`
	Id   uint64 `json:"id"`
}

type ProductRequest struct {
	UserId        uint64   `json:"userid"`
	ProductName   string   `json:"productname"`
	ProductDesc   string   `json:"productdesc"`
	ProductImages []string `json:"productimages"`
	ProductPrice  string   `json:"productprice"`
}
