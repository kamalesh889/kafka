package model

import (
	"time"

	"github.com/lib/pq"
)

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
	ProductImages           pq.StringArray `gorm:"type:text[]"`
	ProductPrice            string
	CompressedProductImages pq.StringArray `gorm:"type:text[]"`
	CreatedAt               time.Time
	UpdatedAt               time.Time
}
