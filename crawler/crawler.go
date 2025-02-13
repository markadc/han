package crawler

import (
	"kss/requests"
	"net/http"
	"time"
)

type Crawler struct {
	Name    string
	proxy   string
	timeout time.Duration
}

func NewCrawler(name string) *Crawler {
	return &Crawler{Name: name}
}

func (c *Crawler) GetProxy() string {
	return c.proxy
}

func (c *Crawler) GetTimeout() time.Duration {
	return c.timeout
}

func (c *Crawler) SetProxy(proxy string) {
	c.proxy = proxy
}

func (c *Crawler) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}

func (c *Crawler) Get(url string, headers, params map[string]string) (*Response, error) {
	req, err := requests.MakeGetRequest(url, headers, params)
	if err != nil {
		return nil, err
	}
	return c.Go(req)
}

func (c *Crawler) Post(api string, headers map[string]string, data []byte) (*Response, error) {
	req, err := requests.MakePostRequest(api, headers, data)
	if err != nil {
		return nil, err
	}
	return c.Go(req)
}

func (c *Crawler) PostForm(api string, headers, formData map[string]string) (*Response, error) {
	req, err := requests.MakePostFormRequest(api, headers, formData)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return c.Go(req)
}

func (c *Crawler) Go(req *http.Request) (*Response, error) {
	if c.proxy == "" {
		client := &http.Client{Timeout: c.timeout}
		return GetResponse(req, client)
	} else {
		client, err := requests.CreateProxyClient(c.proxy, c.timeout)
		if err != nil {
			return nil, err
		}
		return GetResponse(req, client)
	}
}

func GetResponse(req *http.Request, client *http.Client) (*Response, error) {
	body, err := requests.Send(req, client)
	if err != nil {
		return nil, err
	}
	res := &Response{Content: body, Text: string(body)}
	return res, nil
}
