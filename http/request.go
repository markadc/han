package http

import (
	"bytes"
	"encoding/json"
	"han"
	"io"
	"net/http"
)

func Get(url string, headers, params han.S) (*Response, error) {
	// 构造请求
	req, err := http.NewRequest("GET", MakeURL(url, params), nil)
	if err != nil {
		return nil, err
	}

	// 设置请求头
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	// 读取内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 返回响应
	return &Response{Request: req, Text: string(body), StatusCode: resp.StatusCode}, nil
}

// JSON请求体
func Post(URL string, headers, params han.S, payload han.S) (*Response, error) {
	// 请求体
	some, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// 构造请求
	req, err := http.NewRequest("POST", MakeURL(URL, params), bytes.NewBuffer(some))

	if err != nil {
		return nil, err
	}

	// 设置请求头
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	// 读取内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 返回响应
	return &Response{Request: req, Text: string(body), StatusCode: resp.StatusCode}, nil

}

// 请求表单
func PostForm(URL string, headers, params han.S, formData han.S) (*Response, error) {
	// 构造请求
	req, err := http.NewRequest("POST", MakeURL(URL, params), bytes.NewBuffer(FormDataEncode(formData)))
	if err != nil {
		return nil, err
	}

	// 设置请求头
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	// 读取内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 返回响应
	return &Response{Request: req, Text: string(body), StatusCode: resp.StatusCode}, nil

}
