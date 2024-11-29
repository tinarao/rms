package db

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID                 uint           `json:"id" gorm:"primaryKey"`
	AddedDishes        []Dish         `json:"addedDishes" gorm:"foreignKey:CreatorId"`
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
	ID          uint                `json:"id" gorm:"primaryKey"`
	Name        string              `json:"name" gorm:"uniqueIndex"`
	Slug        string              `json:"slug" gorm:"uniqueIndex"`
	Description string              `json:"description"`
	Pictures    []RestaurantPicture `json:"pictures" gorm:"many2many:restaurant_rstpictures;"`
	Creator     Admin               `json:"creator"`
	CreatorId   uint                `json:"creatorId"`
	Orders      []Order             `json:"orders" gorm:"foreignKey:RestaurantID"`
	Dishes      []Dish              `json:"dishes" gorm:"foreignKey:RestaurantID"`
	CreatedAt   time.Time           `json:"createdAt"`
	UpdatedAt   time.Time           `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt      `json:"deletedAt" gorm:"index"`
}

type RestaurantPicture struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"string"`
	Path string `json:"string"`

	RestaurantId uint `json:"restaurantId"`
}

type Dish struct {
	ID             uint       `json:"id" gorm:"primaryKey"`
	Name           string     `json:"name" gorm:"uniqueIndex"`
	Slug           string     `json:"slug" gorm:"uniqueIndex"`
	Description    *string    `json:"description"`
	Tags           []string   `json:"tags" gorm:"type:text[]"`
	PurchasesCount uint       `json:"purchases" gorm:"default:0"`
	IsPublished    bool       `json:"isPublished" gorm:"default:false"`
	Orders         []*Order   `json:"-" gorm:"many2many:dish_orders;"`
	Restaurant     Restaurant `json:"restaurant"`
	RestaurantID   uint       `json:"restaurantId"`
	Creator        User       `json:"creator"`
	CreatorId      uint       `json:"creatorId"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type Order struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	Dishes       []*Dish    `json:"dishes" gorm:"many2many:order_dishes;"`
	Restaurant   Restaurant `json:"restaurant"`
	RestaurantID uint       `json:"restaurantId"`
	Sum          uint       `json:"sum"`
	User         User       `json:"user"`
	UserId       uint       `json:"userId"`
}
