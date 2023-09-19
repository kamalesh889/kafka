package model

import (
	"log"
)

type Repository interface {
	CreateUser(*User) (uint64, error)
	GetUser(userId uint64) error
	CreateProduct(*Product) (uint64, error)
	GetProduct(productid uint64) (*Product, error)
	UpdateProduct(*Product) error
}

func (db *Database) CreateUser(userDetails *User) (uint64, error) {

	result := db.DbConn.Create(userDetails)
	if result.Error != nil {
		log.Println("Error in creating the user", result.Error)
		return 0, result.Error
	}

	userid := userDetails.Id
	return userid, nil
}

func (db *Database) GetUser(userId uint64) error {

	user := User{}

	err := db.DbConn.First(&user, "id=?", userId).Error
	if err != nil {
		log.Println("Error in Fetching user details", err)
		return err
	}

	return nil
}

func (db *Database) CreateProduct(productDetails *Product) (uint64, error) {

	result := db.DbConn.Create(productDetails)
	if result.Error != nil {
		log.Println("Error in creating the product", result.Error)
		return 0, result.Error
	}

	productId := productDetails.ProductId

	return productId, nil
}

func (db *Database) GetProduct(productid uint64) (*Product, error) {

	product := Product{}

	err := db.DbConn.First(&product, "product_id=?", productid).Error
	if err != nil {
		log.Println("Error in Fetching product details", err)
		return nil, err
	}

	return &product, nil

}

func (db *Database) UpdateProduct(updatedProduct *Product) error {

	err := db.DbConn.Model(Product{}).
		Where("product_id = ?", updatedProduct.ProductId).
		Updates(Product{
			CompressedProductImages: updatedProduct.CompressedProductImages,
			UpdatedAt:               updatedProduct.UpdatedAt,
		}).Error

	if err != nil {
		log.Println("Error in Updating product details", err)
		return err
	}

	return nil
}
