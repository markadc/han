package reqs

import (
	"encoding/json"
	"kss"
	"log"
	"net/http"
)

type origin = *http.Response

// 响应
type Response struct {
	origin
	Content []byte
	Text    string
}

// JSON 获取响应的JSON格式数据
func (r Response) JSON() (kss.A, error) {
	var jsonData kss.A
	err := json.Unmarshal(r.Content, &jsonData)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

// JsonStringify JSON字符串格式化，优化打印显示
func (r Response) JsonStringify() string {
	jsonData, err := r.JSON()
	if err != nil {
		panic(err)
	}
	bs, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		log.Fatalf("JSON serialization failed: %v", err)
	}
	return string(bs)
}
