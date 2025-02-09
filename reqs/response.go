package reqs

import (
	"encoding/json"
	"io"
	"kss"
	"log"
	"net/http"
)

// 响应
type Response struct {
	Request    *http.Request
	StatusCode int
	Body       io.ReadCloser
	Content    []byte
	Text       string
}

// 获取响应的JSON格式数据
func (r Response) JSON() (kss.A, error) {
	var jsonData kss.A
	err := json.Unmarshal([]byte(r.Text), &jsonData)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

// JSON字符串格式化，优化打印显示
func (r Response) JsonStringify() string {
	jsonData, err := r.JSON()
	if err != nil {
		panic(err)
	}
	b, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		log.Fatalf("JSON serialization failed: %v", err)
	}
	return string(b)
}
