package repositories

import (
	"echo-go/app/models"
	"echo-go/app/utils"
	"echo-go/config"

	"github.com/google/uuid"
)

func RegisterUser(user models.User) (models.User, error) {
	hash, _ := utils.HashPassword(user.Password)

	input := models.User{ID: uuid.New(), Name: user.Name, Username: user.Username, Password: hash}

	err := config.DB.Create(&input)

	if err.Error != nil {
		return input, err.Error
	}
	return input, nil
}

func Login(user models.User) (models.User, error) {
	data := models.User{Username: user.Username}

	err := config.DB.Find(&data)

	if err.Error != nil {
		return data, err.Error
	}
	return data, nil
}
