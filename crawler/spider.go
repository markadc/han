package crawler

import (
	"kss/requests"
	"net/http"
	"time"
)

type Spider struct {
	Name    string
	proxy   string
	timeout time.Duration
}

func NewSpider(name string) *Spider {
	return &Spider{Name: name}
}

func (s *Spider) GetProxy() string {
	return s.proxy
}

func (s *Spider) GetTimeout() time.Duration {
	return s.timeout
}

func (s *Spider) SetProxy(proxy string) {
	s.proxy = proxy
}

func (s *Spider) SetTimeout(timeout time.Duration) {
	s.timeout = timeout
}

func (s *Spider) Get(url string, headers, params map[string]string) (*Response, error) {
	req, err := requests.MakeGetRequest(url, headers, params)
	if err != nil {
		return nil, err
	}
	return s.Go(req)
}

func (s *Spider) Post(api string, headers map[string]string, data []byte) (*Response, error) {
	req, err := requests.MakePostRequest(api, headers, data)
	if err != nil {
		return nil, err
	}
	return s.Go(req)
}

func (s *Spider) PostForm(api string, headers, formData map[string]string) (*Response, error) {
	req, err := requests.MakePostFormRequest(api, headers, formData)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return s.Go(req)
}

func (s *Spider) Go(req *http.Request) (*Response, error) {
	if s.proxy == "" {
		client := &http.Client{Timeout: s.timeout}
		return s.Send(req, client)
	} else {
		client, err := requests.CreateProxyClient(s.proxy, s.timeout)
		if err != nil {
			return nil, err
		}
		return s.Send(req, client)
	}
}

func (s *Spider) Send(req *http.Request, client *http.Client) (*Response, error) {
	body, err := requests.Send(req, client)
	if err != nil {
		return nil, err
	}
	res := &Response{Content: body, Text: string(body)}
	return res, nil
}
