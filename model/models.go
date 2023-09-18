package model

import "time"

type User struct {
	Id        uint64 `gorm:"primaryKey;autoIncrement"`
	Name      string
	Mobile    string
	Latitude  string
	Longitude string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ProductId               uint64 `gorm:"primaryKey;autoIncrement"`
	ProductName             string
	ProductDesc             string
	ProductImages           []string `gorm:"type:varchar(255)[]"`
	ProductPrice            string
	CompressedProductImages []string `gorm:"type:varchar(255)[]"`
	CreatedAt               time.Time
	UpdatedAt               time.Time
}
