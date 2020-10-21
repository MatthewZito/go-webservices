package domain

import (
	"errors"
	"fmt"
)

var (
	users = map[int64]*User{
		123: {
			ID:        123,
			FirstName: "Fred",
			LastName:  "Flintstone",
			Email:     "fred_flint@bedrock.com",
		},
	}
)

func GetUser(userID int64) (*User, error) {
	if user, ok := users[userID]; ok {
		return user, nil
	} else {
		return nil, errors.New(fmt.Sprintf("[-] User with ID %v not found", userID))
	}
}
