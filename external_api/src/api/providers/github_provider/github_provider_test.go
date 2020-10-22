package github_provider

import "testing"

func TestGetAuthHeader(t *testing.T) {
	header := getAuthHeader("YXV0aHRva2VuCg==")
}
