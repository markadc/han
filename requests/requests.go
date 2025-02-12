package requests

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Get 发送 GET 请求
func Get(url string, headers, params map[string]string) ([]byte, error) {
	query := urlValues(params)
	if query != "" {
		url = fmt.Sprintf("%s?%s", url, query)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req, headers)
	return do(req)

}

// Post 发送 POST 请求（JSON 数据）
func Post(url string, headers map[string]string, data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	setHeaders(req, headers)
	return do(req)
}

// PostForm 发送 POST 请求（表单数据）
func PostForm(url string, headers, formData map[string]string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(urlValues(formData)))
	if err != nil {
		return nil, err
	}
	setHeaders(req, headers)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return do(req)
}

// do 发送请求，获取响应
func do(req *http.Request) ([]byte, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = res.Body.Close() }()
	return io.ReadAll(res.Body)
}

// urlValues 将 map 转换为 URL 查询参数或表单数据
func urlValues(data map[string]string) string {
	values := url.Values{}
	for key, value := range data {
		values.Add(key, value)
	}
	return values.Encode()
}

// setHeaders 设置请求头
func setHeaders(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}
