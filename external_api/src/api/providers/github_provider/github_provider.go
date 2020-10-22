package github_provider

import (
	"api/src/api/domain/github"
	"fmt"
)

const (
	authHeader       = "Authorization"
	authHeaderFormat = "token %s"
)

func getAuthHeader(token string) string {
	return fmt.Sprintf(authHeaderFormat, token)
}

func CreateRepo(token string, github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GitHubErrorResponse) {
	header := getAuthHeader(token)
	return github.CreateRepoResponse{}, github.GitHubErrorResponse{}
}
