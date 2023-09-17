package model

import "time"

type User struct {
	Id        uint64 `gorm:"primaryKey"`
	Name      string
	Mobile    string
	Latitude  string
	Longitude string
	CreatedAt time.Time
	UpdatedAt time.Time
}
