package db

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            uint `json:"id" gorm:"primaryKey"`
	IsAdmin       bool `json:"isAdmin"`
	AddedProducts []Product

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Product struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Creator User
}
