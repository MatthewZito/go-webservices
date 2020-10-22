package domain

import (
	"fmt"
	"net/http"

	"github.com/MatthewZito/go-microservices/mvc/utils"
)

// mock db
var (
	users = map[int64]*User{
		123: {
			ID:        123,
			FirstName: "Fred",
			Surname:   "Flintstone",
			Email:     "fred_flint@bedrock.com",
		},
	}
)

// implement pointers so as to return nil
func GetUser(userID int64) (*User, *utils.ServiceError) {
	if user, ok := users[userID]; ok {
		return user, nil
	} else {
		return nil, &utils.ServiceError{
			Message:    fmt.Sprintf("User with ID %v not found", userID),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
}
