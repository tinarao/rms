package auth

import (
	"context"
	"errors"
	"fmt"
	"rms-api/db"
	"rms-api/redis"

	"github.com/aidarkhanov/nanoid"
	"gorm.io/gorm"
)

func CheckAuthRequest(phone string) error {
	existingRequest := &db.AuthRequest{}
	if result := db.Client.Where("phone = ?", phone).First(&existingRequest); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("auth request for this number does not exist")
		}
	}

	return nil
}

func HandleSecondAuthorizationStep(phone string) error {
	user := &db.User{}
	deletedRequest := &db.AuthRequest{}

	result := db.Client.Where("phone = ?", phone).Delete(&deletedRequest)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("auth request for this number does not exist")
	}

	result = db.Client.Where("phone = ?", phone).First(&user)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		registeredUser := db.User{
			Phone: phone,
		}

		result := db.Client.Create(&registeredUser)
		if result.Error != nil {
			return fmt.Errorf("an error occured when creating an user: %s", result.Error)
		}

		user.ID = registeredUser.ID
	}

	sessionId := nanoid.New()
	redis.Client.Set(context.Background(), sessionId, user.ID, 0)

	db.Client.Where("id = ?", user.ID).Update("sessionId", sessionId)

	return nil
}
