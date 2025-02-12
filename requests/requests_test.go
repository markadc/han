package requests

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestReq(t *testing.T) {
	// 示例：GET 请求
	headers := map[string]string{"User-Agent": "Go requests"}
	params := map[string]string{"name": "wauo", "age": "22"}
	response, _ := Get("https://httpbin.org/get", headers, params)
	fmt.Println("GET\n", string(response))

	// 示例：POST 请求（JSON 数据）
	jsonPayload := map[string]string{"type": "json"}
	bs, _ := json.Marshal(jsonPayload)
	response, _ = Post("https://httpbin.org/post", headers, bs)
	fmt.Println("POST\n:", string(response))

	// 示例：POST 请求（表单数据）
	formData := map[string]string{"type": "form-data"}
	response, _ = PostForm("https://httpbin.org/post", headers, formData)
	fmt.Println("POST Form\n", string(response))
}

type Animal struct {
	Name  string
	Age   int
	Color string
}

func (a Animal) String() string {
	return fmt.Sprintf("%s 的 %s 现在 %d 岁了", a.Color, a.Name, a.Age)
}

func (a Animal) Demo() {
	fmt.Println("哈哈")
}

func TestAnimal(t *testing.T) {
	a := Animal{"小白", 2, "白色"}
	fmt.Println(a)
}
