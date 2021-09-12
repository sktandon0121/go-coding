package services

import (
	"github.com/subodhqss/testing/models"
)

func GetUser(id []int) *models.User {
	if len(id) == 0 {
		return nil
	}
	user := models.User{
		ID:   1,
		Name: "Alam",
	}
	return &user
}
