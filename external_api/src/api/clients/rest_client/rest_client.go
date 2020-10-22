package rest_client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Post(URL string, body interface{}, headers http.Header) (*http.Response, error) {
	JSONBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, URL, bytes.NewReader(JSONBytes))
	if err != nil {
		return nil, err
	}

	request.Header = headers
	client := http.Client{}

}
