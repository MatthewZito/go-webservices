package domain

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "Unexpected user 0.")

	assert.NotNil(t, err, "An error is expected.")

	assert.EqualValues(t, err.StatusCode, http.StatusNotFound, "Expected a 404 status code.")
	assert.EqualValues(t, err.Code, "not_found", "Expected a code of 'not_found'.")
	assert.EqualValues(t, err.Message, "User with ID 0 not found", "Unexpected error message.")
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.NotNil(t, user, "Expected a user JSON object.")
	assert.Nil(t, err, "Struct ServiceError should be nil.")

	var expected User = User{
		ID:        123,
		FirstName: "Fred",
		Surname:   "Flintstone",
		Email:     "fred_flint@bedrock.com",
	}

	assert.EqualValues(t, user.ID, expected.ID, fmt.Sprintf("Expected ID to be %d.", expected.ID))
	assert.EqualValues(t, user.FirstName, expected.FirstName, fmt.Sprintf("Expected FirstName to be '%s'.", expected.FirstName))
	assert.EqualValues(t, user.Surname, expected.Surname, fmt.Sprintf("Expected Surname to be '%s'.", expected.Surname))
	assert.EqualValues(t, user.Email, expected.Email, fmt.Sprintf("Expected Email to be '%s'.", expected.Email))
}
