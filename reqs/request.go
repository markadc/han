package reqs

import (
	"bytes"
	"encoding/json"
	"han"
	"net/http"
)

func Get(url string, headers, params han.S) (*Response, error) {
	req, err := http.NewRequest("GET", MakeURL(url, params), nil)
	if err != nil {
		return nil, err
	}
	SetHeaders(req, headers)
	client := &http.Client{}
	return Done(client, req)
}

// JSON请求体
func Post(URL string, headers, params han.S, payload han.S) (*Response, error) {
	some, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", MakeURL(URL, params), bytes.NewBuffer(some))
	if err != nil {
		return nil, err
	}
	SetHeaders(req, headers)
	client := &http.Client{}
	return Done(client, req)
}

// 请求表单
func PostForm(URL string, headers, params han.S, formData han.S) (*Response, error) {
	req, err := http.NewRequest("POST", MakeURL(URL, params), bytes.NewBuffer(FormDataEncode(formData)))
	if err != nil {
		return nil, err
	}
	SetHeaders(req, headers)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	return Done(client, req)
}
