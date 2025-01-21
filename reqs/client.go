package reqs

import (
	"bytes"
	"encoding/json"
	"kss"
	"net/http"
	"time"
)

// RequestConfigs 请求配置
type RequestConfigs struct {
	Params kss.S
	Header kss.S
	Proxy  string
}

// Client 客户端
type Client struct {
	Timeout time.Duration
	Proxy   string
}

func NewClient(timeout time.Duration, proxy string) *Client {
	return &Client{Timeout: timeout, Proxy: proxy}
}

// GetClient 代理和超时
func (c *Client) GetClient() (*http.Client, error) {
	transport, err := MakeProxyTransport(c.Proxy)
	if err != nil {
		return nil, err
	}
	client := &http.Client{Timeout: c.Timeout, Transport: transport}
	return client, nil
}

// Do 发送请求
func (c *Client) Do(req *http.Request, headers kss.S) (*Response, error) {
	ReqSetHeaders(req, headers)
	client, err := c.GetClient()
	if err != nil {
		return nil, err
	}
	return Done(client, req)
}

// Get GET请求
func (c *Client) Get(url string, headers, params kss.S) (*Response, error) {
	req, err := http.NewRequest("GET", MakeURL(url, params), nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req, headers)
}

// Post POST请求（JSON请求体）
func (c *Client) Post(URL string, headers, params kss.S, payload kss.S) (*Response, error) {
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

// PostForm POST（请求表单）
func (c *Client) PostForm(URL string, headers, params kss.S, formData kss.S) (*Response, error) {
	req, err := http.NewRequest("POST", MakeURL(URL, params), bytes.NewBuffer(FormDataEncode(formData)))
	if err != nil {
		return nil, err
	}
	delete(headers, "Content-Type")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return c.Do(req, headers)
}
