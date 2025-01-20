package http

import (
	"han"
	"io"
	"net/http"
	"time"
)

type RequestConfigs struct {
	Params han.S
	Header han.S
	Proxy  string
}

type HTTPClient struct {
	Timeout time.Duration
	Proxy   string
}

func NewHTTPClient(timeout time.Duration, proxy string) *HTTPClient {
	return &HTTPClient{Timeout: timeout, Proxy: proxy}
}

func (h *HTTPClient) Get(url string, headers, params han.S) (*Response, error) {
	req, err := http.NewRequest("GET", MakeURL(url, params), nil)
	if err != nil {
		return nil, err
	}

	// 设置请求头
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// 发送请求
	tr, err := MakeProxy(h.Proxy)
	if err != nil {
		return nil, err
	}
	client := &http.Client{Timeout: h.Timeout, Transport: tr}
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
