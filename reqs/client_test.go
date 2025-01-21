package reqs

import (
	"fmt"
	"han"
	"testing"
	"time"
)

func TestNewClient_Get(t *testing.T) {
	cli := NewClient(500*time.Millisecond, "", nil)
	res, err := cli.Get("https://www.baidu.com", nil, nil)
	if err != nil {
		t.Fatalf("失败了 %s", err)
	}
	fmt.Printf("你的网速还可以 %v %v\n", res.Request.URL, res.StatusCode)
}

func TestClient_UseProxy(t *testing.T) {
	srcURL := "https://github.com"
	cli := NewClient(5*time.Second, "http://localhost:10809", nil)
	res, err := cli.Get(srcURL, nil, nil)
	if err != nil {
		t.Fatalf("或者代理有误 %s", err)
	}
	fmt.Printf("%v 响应长度 %v\n", res.Request.URL, res.StatusCode)
}

func TestClient(t *testing.T) {
	cli := &Client{Headers: han.S{"User-Agent": GenRandomUA()}}
	link := "https://www.baidu.com"
	resp, err := cli.Get(link, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Request.URL, resp.Request.Header)
}
