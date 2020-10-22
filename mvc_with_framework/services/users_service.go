package services

import (
	"github.com/MatthewZito/go-microservices/mvc/domain"
	"github.com/MatthewZito/go-microservices/mvc/utils"
)

// use pointer so as to be able to assign nil
func GetUser(userID int64) (*domain.User, *utils.ServiceError) {
	return domain.GetUser(userID)
}
