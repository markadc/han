package reqs

import (
	"encoding/json"
	"han"
	"log"
	"net/http"
)

// 响应
type Response struct {
	StatusCode int
	Text       string
	Request    *http.Request
}

// 获取响应的JSON格式数据
func (r Response) JSON() (han.A, error) {
	var jsonData han.A
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
