package crawler

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSpider(t *testing.T) {
	s := NewSpider("test")

	// GET 请求
	headers := map[string]string{"User-Agent": "Go requests"}
	params := map[string]string{"name": "wauo", "age": "22"}
	res, _ := s.Get("https://httpbin.org/get", headers, params)
	fmt.Println("GET\n", res.Text)

	// POST 请求（请求体）
	jsonPayload := map[string]string{"msg": "RequestBody => Success"}
	bs, _ := json.Marshal(jsonPayload)
	res, _ = s.Post("https://httpbin.org/post", headers, bs)
	fmt.Println("POST\n", res.Text)

	// POST 请求（表单数据）
	formData := map[string]string{"msg": " FormData => Success"}
	res, _ = s.PostForm("https://httpbin.org/post", headers, formData)
	fmt.Println("POST Form\n", res.Text)
}
