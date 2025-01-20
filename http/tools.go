package http

import (
	"fmt"
	"han"
	"net/http"
	"net/url"
	"strings"
)

// 构造完整的请求URL
func MakeURL(baseURL string, params han.S) string {
	if !strings.Contains(baseURL, "?") {
		baseURL += "?"
	} else {
		baseURL += "&"
	}
	for key, value := range params {
		baseURL += fmt.Sprintf("%s=%s&", key, value)
	}
	return baseURL[:len(baseURL)-1]
}

// 对请求表单进行编码
func FormDataEncode(formData han.S) []byte {
	encoded := make([]byte, 0)
	for key, value := range formData {
		encoded = append(encoded, []byte(fmt.Sprintf("%s=%s&", key, value))...)
	}
	encoded = encoded[:len(encoded)-1]
	return encoded
}

// 获取 http https 的字符串
func GetProtocols(protocol string) (string, string) {
	var p1, p2 string
	if strings.HasSuffix(protocol, "https") {
		p1 = "http" + strings.TrimPrefix(protocol, "https")
		p2 = protocol
	} else {
		p1 = protocol
		p2 = "https" + strings.TrimPrefix(protocol, "http")
	}
	return p1, p2
}

// 设置代理
func MakeProxy(proxy string) (*http.Transport, error) {
	if proxy == "" {
		return &http.Transport{}, nil
	}
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return nil, err
	}
	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	return transport, nil
}
