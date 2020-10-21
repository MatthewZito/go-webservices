package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/MatthewZito/go-microservices/mvc/services"
)

// GetUser processes GET requests for user_id <int>
func GetUser(resp http.ResponseWriter, req *http.Request) {
	userID, err := (strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64))
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("[-] user_id must be an integer"))
		return
	}
	log.Printf("[+] Processing request for user_id %v", userID)
	user, err := services.GetUser(userID)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(fmt.Sprintf("[-] An error occurred. See: %v", err.Error())))
		return
	}

	json, _ := json.Marshal(user)
	resp.Write(json)

}
