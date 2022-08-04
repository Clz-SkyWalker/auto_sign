package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

func NewUtilsRequest(param UtilsRequestParam) *UtilsRequest {
	return &UtilsRequest{
		UtilsRequestParam: param,
		client:            &http.Client{},
	}
}

// 请求工具类对象
type UtilsRequestParam struct {
	Url    string
	Body   interface{}
	Params map[string]string
	Header http.Header
}

type UtilsRequest struct {
	client *http.Client
	UtilsRequestParam
	Respose *http.Response
	Request *http.Request
	Err     error
}

func (r *UtilsRequest) Get() {
	uri, err := url.Parse(r.Url)
	if err != nil {
		r.Err = err
		return
	}
	urlRow := url.Values{}
	for key, value := range r.Params {
		urlRow.Add(key, value)
	}
	uri.RawQuery = urlRow.Encode()
	r.Request, r.Err = http.NewRequest(http.MethodGet, uri.String(), nil)
	if r.Err != nil {
		return
	}

	r.Request.Header = r.Header
	r.client.Timeout = time.Duration(time.Second * 10)
	r.Respose, r.Err = r.client.Do(r.Request)
}

func (r *UtilsRequest) Post() {
	uri, err := url.Parse(r.Url)
	if err != nil {
		r.Err = err
		return
	}
	urlRow := url.Values{}
	for key, value := range r.Params {
		urlRow.Add(key, value)
	}
	uri.RawQuery = urlRow.Encode()

	var byteList []byte
	byteList, r.Err = json.Marshal(r.Body)
	if r.Err != nil {
		return
	}
	r.Request, r.Err = http.NewRequest(http.MethodPost, uri.String(), bytes.NewReader(byteList))
	if r.Err != nil {
		return
	}

	r.Request.Header = r.Header
	r.client.Timeout = time.Duration(time.Second * 10)
	r.Respose, r.Err = r.client.Do(r.Request)
}
