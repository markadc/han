package http

import (
	"fmt"
	"math/rand"
	"time"
)

// 使用全局变量初始化 rand
var r *rand.Rand

// 初始化 rand 实例
func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano())) // 使用新的方式生成随机数生成器
}

// 生成随机操作系统
func genRandomOS() string {
	osChoices := []string{
		"Windows NT 10.0; Win64; x64",
		"Windows NT 6.1; WOW64",
		"Macintosh; Intel Mac OS X 10_15_6",
		"Linux x86_64",
		"Windows NT 6.3; Trident/7.0",
	}
	return osChoices[r.Intn(len(osChoices))]
}

// 生成随机浏览器
func genRandomBrowser() string {
	browserChoices := []struct {
		Name    string
		Version string
	}{
		{"Chrome", "70"},
		{"Firefox", "31"},
		{"Edge", "31"},
		{"Safari", "537.36"},
		{"Opera", "70"},
	}
	browser := browserChoices[r.Intn(len(browserChoices))]
	return fmt.Sprintf("%s/%s", browser.Name, browser.Version)
}

// 生成随机 User-Agent
func GenRandomUA() string {
	os := genRandomOS()
	browser := genRandomBrowser()
	return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) %s Safari/537.36", os, browser)
}
