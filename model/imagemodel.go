package model

import "time"

type Image struct {
	ProductId               uint64 `gorm:"primaryKey;autoIncrement"`
	ProductName             string
	ProductDesc             string
	ProductImages           []string `gorm:"type:varchar(255)[]"`
	ProductPrice            string
	CompressedProductImages []string `gorm:"type:varchar(255)[]"`
	CreatedAt               time.Time
	UpdatedAt               time.Time
}
