package reqs

import (
	"bytes"
	"encoding/json"
	"kss"
	"net/http"
)

// Get GET请求
func Get(url string, headers, params kss.S) (*Response, error) {
	req, err := http.NewRequest("GET", MakeURL(url, params), nil)
	if err != nil {
		return nil, err
	}
	return Do(req, headers)
}

// Post POST请求（JSON请求体）
func Post(URL string, headers, params kss.S, data kss.S) (*Response, error) {
	some, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", MakeURL(URL, params), bytes.NewBuffer(some))
	if err != nil {
		return nil, err
	}
	return Do(req, headers)
}

// PostForm POST请求（请求表单）
func PostForm(URL string, headers, params kss.S, formData kss.S) (*Response, error) {
	req, err := http.NewRequest("POST", MakeURL(URL, params), bytes.NewBuffer(FormDataEncode(formData)))
	if err != nil {
		return nil, err
	}
	ReqSetHeaders(req, headers)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return Do(req, headers)
}
