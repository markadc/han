package requests

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestReq(t *testing.T) {
	// GET 请求
	headers := map[string]string{"User-Agent": "Go requests"}
	params := map[string]string{"name": "wauo", "age": "22"}
	body, _ := Get("https://httpbin.org/get", headers, params)
	fmt.Println("GET\n", string(body))

	// POST 请求（JSON 数据）
	jsonPayload := map[string]string{"type": "json"}
	bs, _ := json.Marshal(jsonPayload)
	body, _ = Post("https://httpbin.org/post", headers, bs)
	fmt.Println("POST\n:", string(body))

	// POST 请求（表单数据）
	formData := map[string]string{"type": "form-data"}
	body, _ = PostForm("https://httpbin.org/post", headers, formData)
	fmt.Println("POST Form\n", string(body))
}

func TestCreateUrl(t *testing.T) {
	url := "https://httpbin.org/get"
	params := map[string]string{"name": "wauo", "age": "22"}
	newUrl := CreateUrl(url, params)
	fmt.Println(newUrl)

	url2 := "https://httpbin.org/get?age=18"
	params2 := map[string]string{"name": "wauo", "age": "22"}
	newUrl2 := CreateUrl(url2, params2)
	fmt.Println(newUrl2)
}

func TestCreateProxyClient(t *testing.T) {
	client, err := CreateProxyClient("http://localhost:10809", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(client)
}
