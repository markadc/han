package requests

import (
	"bytes"
	"net/http"
	"strings"
)

func MakeGetRequest(url string, headers, params map[string]string) (*http.Request, error) {
	req, err := http.NewRequest("GET", CreateUrl(url, params), nil)
	if err != nil {
		return nil, err
	}
	SetHeaders(req, headers)
	return req, nil
}

func MakePostRequest(api string, headers map[string]string, data []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", api, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	SetHeaders(req, headers)
	return req, nil
}

func MakePostFormRequest(api string, headers, formData map[string]string) (*http.Request, error) {
	req, err := http.NewRequest("POST", api, strings.NewReader(UrlValues(formData)))
	if err != nil {
		return nil, err
	}
	SetHeaders(req, headers)
	return req, nil
}
