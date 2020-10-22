package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/MatthewZito/go-microservices/mvc/services"
	"github.com/MatthewZito/go-microservices/mvc/utils"
)

// GetUser processes GET requests for user_id <int>
func GetUser(resp http.ResponseWriter, req *http.Request) {
	userID, err := (strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64))

	if err != nil {
		userErr := &utils.ServiceError{
			Message:    "user_id must be an integer",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		resp.WriteHeader(userErr.StatusCode)

		jsonVal, _ := json.Marshal(userErr)
		resp.Write(jsonVal)
		return
	}

	log.Printf("[+] Processing request for user_id %v", userID)

	user, userErr := services.GetUser(userID)

	if userErr != nil {
		jsonVal, _ := json.Marshal(userErr)
		resp.Write(jsonVal)
		return
	}

	jsonVal, _ := json.Marshal(user)
	resp.Write(jsonVal)
}
