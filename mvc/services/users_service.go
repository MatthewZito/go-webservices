package services

import (
	"github.com/MatthewZito/go-microservices/mvc/domain"
)

func GetUser(userID int64) (*domain.User, error) {
	return domain.GetUser(userID)
}
