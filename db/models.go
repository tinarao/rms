package db

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID                 uint           `json:"id" gorm:"primaryKey"`
	AddedProducts      []Product      `json:"addedProducts" gorm:"foreignKey:CreatorId"`
	CreatedRestaurants []Restaurant   `json:"createdRestaurants" gorm:"foreignKey:CreatorId"`
	LatestIP           *string        `json:"latestIp"`
	FirstName          string         `json:"firstName"`
	LastName           string         `json:"lastName"`
	Email              string         `json:"email" gorm:"uniqueIndex"`
	Password           string         `json:"-"`
	CreatedAt          time.Time      `json:"createdAt"`
	UpdatedAt          time.Time      `json:"updatedAt"`
	DeletedAt          gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	SessionId *string        `json:"sessionId"`
	Orders    []Order        `json:"orders" gorm:"foreignKey:UserId"`
	Phone     string         `json:"phone" gorm:"uniqueIndex"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type AuthRequest struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Phone     string    `json:"phone" gorm:"uniqueIndex"`
	CreatedAt time.Time `json:"createdAt"`
}

type Restaurant struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" gorm:"uniqueIndex"`
	Slug        string  `json:"slug" gorm:"uniqueIndex"`
	Description string  `json:"description"`
	Orders      []Order `json:"orders" gorm:"foreignKey:RestaurantID"`

	Pictures []string `json:"pictures" gorm:"type:text[]"`

	Creator   Admin `json:"creator"`
	CreatorId uint  `json:"creatorId"`
}

type Product struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	SessionId *string  `json:"sessionId"`
	Purchases uint     `json:"purchases"`
	Orders    []*Order `json:"orders" gorm:"many2many:product_orders;"`
	Creator   User     `json:"creator"`
	CreatorId uint     `json:"creatorId"`
}

type Order struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	Products     []*Product `json:"products" gorm:"many2many:user_languages;"`
	Restaurant   Restaurant `json:"restaurant"`
	RestaurantID uint       `json:"restaurantId"`
	Sum          uint       `json:"sum"`
	User         User       `json:"user"`
	UserId       uint       `json:"userId"`
}
