package repositories

import (
	"echo-go/app/models"
	"echo-go/app/utils"
	"echo-go/config"

	"github.com/google/uuid"
)

func GetUser(user []models.User) ([]models.User, error) {
	data := config.DB.Find(&user)

	if data.Error != nil {
		return user, data.Error
	}

	return user, nil

}

func CreateUser(user models.User) (models.User, error) {
	// hash, _ := utils.HashPassword(user.Password)

	input := models.User{ID: uuid.New(), Name: user.Name, Username: user.Username, Password: user.Password}

	err := config.DB.Create(&input)

	if err.Error != nil {
		return input, err.Error
	}
	return input, nil
	// db := config.GetDB
	// sqlStatement := `INSERT INTO users (id, username, password) VALUES ($1, $2, $3) RETURNING id`
	// err := db.QueryRow(sqlStatement, uuid.New(), user.Username, user.Password).Scan(&user.ID)
	// if err != nil {
	// 	return user, err
	// }
	// return user, nil

}

func GetUserByID(user models.User, id uuid.UUID) (models.User, error) {
	data := models.User{ID: id}

	err := config.DB.Find(&data)

	if err.Error != nil {
		return data, err.Error
	}
	return data, nil
}

func UpdateUser(user models.User, id uuid.UUID) (models.User, error) {
	data := models.User{ID: id}
	hash, _ := utils.HashPassword(user.Password)

	input := models.User{Name: user.Name, Username: user.Username, Password: hash}

	err := config.DB.Where("id = ?", id).Updates(&input)

	if err.Error != nil {
		return data, err.Error
	}
	return data, nil
}

func DeleteUser(user models.User, id uuid.UUID) (models.User, error) {
	data := models.User{ID: id}

	err := config.DB.Delete(&data, id)

	if err.Error != nil {
		return data, err.Error
	}
	return data, nil
}
