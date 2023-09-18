package model

import "log"

type Repository interface {
	CreateUser(*User) (uint64, error)
	GetUser(userId uint64) error
	CreateProduct(*Product) (uint64, error)
}

func (db *Database) CreateUser(userDetails *User) (uint64, error) {

	result := db.DbConn.Create(userDetails)
	if result.Error != nil {
		log.Println("Error in creating the user", result.Error)
	}

	userid := userDetails.Id
	return userid, nil
}

func (db *Database) GetUser(userId uint64) error {

	user := User{}

	err := db.DbConn.First(&user, "id=?", userId).Error
	if err != nil {
		log.Println("Error in Fetching user details", err)
	}

	return nil
}

func (db *Database) CreateProduct(productDetails *Product) (uint64, error) {

	result := db.DbConn.Create(productDetails)
	if result.Error != nil {
		log.Println("Error in creating the product", result.Error)
	}

	productId := productDetails.ProductId

	return productId, nil
}
