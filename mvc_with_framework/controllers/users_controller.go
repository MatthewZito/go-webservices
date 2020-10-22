package controllers

import (
	"net/http"
	"strconv"

	"github.com/MatthewZito/go-microservices/mvc/services"
	"github.com/MatthewZito/go-microservices/mvc/utils"
	"github.com/gin-gonic/gin"
)

// GetUser processes GET requests for user_id <int>
func GetUser(c *gin.Context) {
	userID, err := (strconv.ParseInt(c.Param("user_id"), 10, 64))

	if err != nil {
		userErr := &utils.ServiceError{
			Message:    "user_id must be an integer",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		utils.RespondError(c, userErr)
		return
	}

	user, userErr := services.GetUser(userID)

	if userErr != nil {
		utils.RespondError(c, userErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}
