// 下载表情包

package main

import (
	"fmt"
	"io"
	"kss"
	"kss/loger"
	"kss/reqs"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

type Task = struct {
	Name string
	Link string
}

func main() {
	var keyword string
	var page int

	if len(os.Args) == 3 {
		keyword = os.Args[1]
		pageStr := os.Args[2]
		temp, err := strconv.Atoi(pageStr)
		if err != nil {
			loger.Error("页数不能是 {} ", pageStr)
		}
		page = temp
	} else if len(os.Args) == 2 {
		keyword = os.Args[1]
		page = 1
	} else {
		println("请提供关键词\ngo run <this file> <关键词>")
		return
	}

	// 创建保存图片的目录
	saveDir := fmt.Sprintf("emos/%s", keyword)
	kss.CreatDir(saveDir)

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0",
		"Referer":    "https://www.doutub.com",
	}

	// 获取图片数据
	tasks := []Task{}
	for i := 0; i < page; i++ {
		url := fmt.Sprintf("https://www.doutub.com/search/%s/%d", keyword, i+1)
		res, err := reqs.Get(url, headers, nil)
		if err != nil {
			loger.Error(fmt.Sprintf("%v %v", err, url))
			continue
		}
		re := regexp.MustCompile(`<img alt="(.+?)" data-src="(.+?)" data`)
		matches := re.FindAllStringSubmatch(res.Text, -1)
		if len(matches) == 0 {
			loger.Error("正则没有匹配到结果 {}", url)
			continue
		}
		for _, match := range matches[1:] {
			task := Task{Name: match[1], Link: match[2]}
			tasks = append(tasks, task)
		}
	}

	// 下载图片
	download := func(task Task) {
		defer wg.Done()

		// 创建图片文件
		filePath := fmt.Sprintf("%s/%s.jpg", saveDir, task.Name)
		file, err := os.Create(filePath)
		if err != nil {
			loger.Error("文件 {} 创建失败", filePath)
			return
		}
		defer func() { _ = file.Close() }()

		// 下载图片到本地
		req, _ := http.NewRequest("GET", task.Link, nil)
		for k, v := range headers {
			req.Header.Set(k, v)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			loger.Error("{} {}", err, task.Link)
			return
		}
		n, err := io.Copy(file, res.Body)
		if err != nil {
			loger.Error("下载文件出错：", err)
		}
		loger.Success("文件《{}》下载成功 {}字节", task.Name, n)
	}

	for _, task := range tasks {
		wg.Add(1)
		go download(task)
	}

	wg.Wait()
}
