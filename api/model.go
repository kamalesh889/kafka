package api

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DbConn *gorm.DB
}

func InitializeDB() (*Database, error) {

	// These configs should be coming from enviroment variable , Here i have hard coded for the demonstration

	db, err := gorm.Open(postgres.Open("postgres://postgres:kamalesh@localhost:5432/image"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{db}, nil
}
