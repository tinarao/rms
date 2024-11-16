package db

import (
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	AddedProducts []Product      `json:"addedProducts" gorm:"foreignKey:CreatorId"`
	Phone         string         `json:"email" gorm:"uniqueIndex"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Orders    []Order        `json:"orders" gorm:"foreignKey:UserId"`
	Phone     string         `json:"phone" gorm:"uniqueIndex"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type Product struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	Purchases uint     `json:"purchases"`
	Orders    []*Order `json:"orders" gorm:"many2many:product_orders;"`
	Creator   User     `json:"creator"`
	CreatorId uint     `json:"creatorId"`
}

type Order struct {
	ID       uint       `json:"id" gorm:"primaryKey"`
	Products []*Product `json:"products" gorm:"many2many:user_languages;"`
	Sum      uint       `json:"sum"`
	User     User       `json:"user"`
	UserId   uint       `json:"userId"`
}
