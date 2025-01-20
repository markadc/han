package http

import (
	"fmt"
	"testing"
	"time"
)

func TestNewHTTPClient(t *testing.T) {
	cli := NewHTTPClient(500*time.Millisecond, "")
	res, err := cli.Get("https://www.baidu.com", nil, nil)
	if err != nil {
		t.Fatalf("失败了 %s", err)
	}
	fmt.Printf("你的网速还可以 %v %v\n", res.Request.URL, res.StatusCode)
}

func TestHTTPClient_Get(t *testing.T) {
	srcURL := "https://github.com"
	cli := NewHTTPClient(5*time.Second, "http://localhost:10809")
	res, err := cli.Get(srcURL, nil, nil)
	if err != nil {
		t.Fatalf("或者代理有误 %s", err)
	}
	fmt.Printf("%v 响应长度 %v\n", res.Request.URL, res.StatusCode)
}
