# 项目说明

- 像`python`的`requests`一样发送请求


```go
package main

import (
	"fmt"
	"kss/reqs"
)

func get() {
	url := "https://httpbin.org/get"
	headers := map[string]string{"Name": "KSS-GET"}
	params := map[string]string{"name": "kssgo"}
	res, _ := reqs.Get(url, headers, params)
	fmt.Println(res.JsonStringify())
}

func post() {
	api := "https://httpbin.org/post"
	headers := map[string]string{"Name": "KSS-POST"}
	payload := map[string]string{"ReqType": "JSON"}
	res, _ := reqs.Post(api, headers, nil, payload)
	fmt.Println(res.JsonStringify())
}

func postForm() {
	api := "https://httpbin.org/post"
	headers := map[string]string{"Name": "KSS-POST"}
	formData := map[string]string{"ReqType": "FormData"}
	res, _ := reqs.PostForm(api, headers, nil, formData)
	fmt.Println(res.JsonStringify())
}

func main() {
	get()
	post()
	postForm()
}

```