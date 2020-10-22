package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJSON(t *testing.T) {
	const responseMock = `{"name":"learning_golang","description":"learning golang by writing microservices","homepage":"https://github.com","private":true,"has_issues":false,"has_projects":false,"has_wiki":false}`

	requestMock := CreateRepoRequest{
		Name:        "learning_golang",
		Description: "learning golang by writing microservices",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   false,
		HasProjects: false,
		HasWiki:     false,
	}

	bytes, err := json.Marshal(requestMock)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	assert.EqualValues(t, string(bytes), responseMock)

	var target CreateRepoRequest

	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)

	assert.EqualValues(t, target.Name, requestMock.Name)
	assert.EqualValues(t, target.HasIssues, requestMock.HasIssues)
}
