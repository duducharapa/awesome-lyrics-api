package repositories

import (
	"duducharapa/lyrics/api/dtos"
	"duducharapa/lyrics/api/models"
	"duducharapa/lyrics/database"
)

func ConvertUser(createDTO *dtos.CreateUserDTO, user *models.User) {
	user.Email = createDTO.Email
	user.Password = createDTO.Password
}

func CreateUser(createDTO *dtos.CreateUserDTO, user *models.User) error {
	ConvertUser(createDTO, user)
	err := database.DB.Create(user).Error

	return err
}

func FindUsers(users *[]models.User) error {
	err := database.DB.Find(users).Error

	return err
}
