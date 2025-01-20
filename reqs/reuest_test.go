package reqs

import (
	"encoding/json"
	"fmt"
	"han"
	"testing"
)

func TestGet(t *testing.T) {
	url := "https://cn.bing.com/search"
	headers := han.S{"Crawler": "wauo", "Cookie": "626"}
	params := han.S{"q": "wauo"}
	resp, err := Get(url, headers, params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Request.URL)
	fmt.Println(resp.Request.Header)
	fmt.Println(resp.StatusCode, len(resp.Text))
}

func TestPost(t *testing.T) {
	url1 := "http://localhost:3333/f"
	url2 := "http://localhost:3333/j"

	some := []int{1, 2, 3}
	b, _ := json.Marshal(some)
	info := string(b)

	payload := han.S{"name": "golang", "mark": info}
	res, _ := PostForm(url1, nil, nil, payload)

	fmt.Println("FormData")
	fmt.Println(res.JsonStringify())
	payload = han.S{"name": "golang", "mark": info}
	res, _ = Post(url2, nil, nil, payload)
	fmt.Println("JSON")
	fmt.Println(res.JsonStringify())
}
