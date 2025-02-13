package requests

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Get 发送 GET 请求
func Get(url string, headers, params map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", CreateUrl(url, params), nil)
	if err != nil {
		return nil, err
	}
	SetHeaders(req, headers)
	return DefaultSend(req)
}

// Post 发送 POST 请求（JSON 数据）
func Post(url string, headers map[string]string, data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	SetHeaders(req, headers)
	return DefaultSend(req)
}

// PostForm 发送 POST 请求（表单数据）
func PostForm(url string, headers, formData map[string]string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(UrlValues(formData)))
	if err != nil {
		return nil, err
	}
	SetHeaders(req, headers)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return DefaultSend(req)
}

// DefaultSend 发送请求（默认客户端）
func DefaultSend(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	return Send(req, client)
}

// Send 发送请求（自定义客户端）
func Send(req *http.Request, client *http.Client) ([]byte, error) {
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = res.Body.Close() }()
	return io.ReadAll(res.Body)
}

// UrlValues 将 map 转换为 URL 查询参数或表单数据
func UrlValues(data map[string]string) string {
	values := url.Values{}
	for key, value := range data {
		values.Add(key, value)
	}
	return values.Encode()
}

// SetHeaders 设置请求头
func SetHeaders(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

// CreateProxyClient 创建代理客户端
func CreateProxyClient(proxy string, timeout time.Duration) (*http.Client, error) {
	if proxy == "" {
		return nil, fmt.Errorf("proxy URL cannot be empty")
	}
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return nil, fmt.Errorf("failed to parse proxy URL: %w", err)
	}

	if timeout == 0 {
		timeout = 30 * time.Second
	}

	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.Proxy = http.ProxyURL(proxyURL)

	client := &http.Client{Timeout: timeout, Transport: transport}
	return client, nil
}

// CreateUrl 构造完整的 url 地址
func CreateUrl(url string, params map[string]string) string {
	query := UrlValues(params)
	if query != "" {
		if strings.Contains(url, "?") {
			url = fmt.Sprintf("%s&%s", url, query)
		} else {
			url = fmt.Sprintf("%s?%s", url, query)
		}
	}
	return url
}
