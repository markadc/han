package reqs

import (
	"bytes"
	"encoding/json"
	"han"
	"net/http"
	"time"
)

// ...
type RequestConfigs struct {
	Params han.S
	Header han.S
	Proxy  string
}

// 客户端
type Client struct {
	Timeout time.Duration
	Proxy   string
	Headers han.S
}

func NewClient(timeout time.Duration, proxy string, headers han.S) *Client {
	return &Client{Timeout: timeout, Proxy: proxy, Headers: headers}
}

// 确认请求头的设置
func (c *Client) ConfirmHeaders(req *http.Request, headers han.S) {
	if len(c.Headers) == 0 {
		SetHeaders(req, headers)
	} else {
		SetHeaders(req, c.Headers)
	}
}

// 代理和超时
func (c *Client) GetClient() (*http.Client, error) {
	transport, err := MakeProxyTransport(c.Proxy)
	if err != nil {
		return nil, err
	}
	client := &http.Client{Timeout: c.Timeout, Transport: transport}
	return client, nil
}

func (c *Client) Do(req *http.Request, headers han.S) (*Response, error) {
	c.ConfirmHeaders(req, headers)
	client, err := c.GetClient()
	if err != nil {
		return nil, err
	}
	return Done(client, req)
}

// GET
func (c *Client) Get(url string, headers, params han.S) (*Response, error) {
	req, err := http.NewRequest("GET", MakeURL(url, params), nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req, headers)
}

// POST JSON请求体
func (c *Client) Post(URL string, headers, params han.S, payload han.S) (*Response, error) {
	some, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", MakeURL(URL, params), bytes.NewBuffer(some))
	if err != nil {
		return nil, err
	}
	return c.Do(req, headers)
}

// POST 请求表单
func (c *Client) PostForm(URL string, headers, params han.S, formData han.S) (*Response, error) {
	req, err := http.NewRequest("POST", MakeURL(URL, params), bytes.NewBuffer(FormDataEncode(formData)))
	if err != nil {
		return nil, err
	}
	c.ConfirmHeaders(req, headers)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client, err := c.GetClient()
	if err != nil {
		return nil, err
	}
	return Done(client, req)
}
