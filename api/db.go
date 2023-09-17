package api

import (
	"Kafka/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() (*model.Database, error) {

	// These configs should be coming from enviroment variable , Here i have hard coded for the demonstration

	db, err := gorm.Open(postgres.Open("postgres://postgres:kamalesh@localhost:5432/image"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// There should be a variable in config to check that we need to migrate or not

	isMigrate := true // Here i am taking default value as true
	if isMigrate {
		Migrate(db)
	}

	return &model.Database{DbConn: db}, nil
}

func Migrate(db *gorm.DB) {

	err := db.AutoMigrate(
		&model.User{}, &model.Image{},
	)

	if err != nil {
		log.Fatalln(err)
		return
	}
}
