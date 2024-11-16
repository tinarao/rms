package auth

import (
	"errors"
	"gorm.io/gorm"
	"rms-api/db"
	"rms-api/dto"
)

func Auth(dto dto.AuthDto) error {
	user := &db.User{}
	result := db.Client.Where("phone = ?", dto.Phone).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		//	Create Profile
	}

	//	tg send code
	//  verify
}
