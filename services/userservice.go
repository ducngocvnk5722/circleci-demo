package services

import (
	"github.com/ducngocvnk57/circleci-demo/models"
)

func GetAllUser() ([]models.User, error) {
	var users []models.User
	err := db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}
