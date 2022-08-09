package utils

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap/zapcore"
)

func NewUtilsRequest(param UtilsRequestParam) *UtilsRequest {
	return &UtilsRequest{
		UtilsRequestParam: param,
		client:            &http.Client{},
	}
}

// 请求工具类对象
type UtilsRequestParam struct {
	Url        string
	Body       interface{}
	Params     map[string]string
	Header     http.Header
	CaCertPath string
	CertFile   string
	KeyFile    string
}

type UtilsRequest struct {
	client *http.Client
	UtilsRequestParam
	Respose *http.Response
	Request *http.Request
	Err     Errno
}

func (r *UtilsRequest) Get() {
	uri, err := url.Parse(r.Url)
	if err != nil {
		r.Err = ErrReqParse.WithErr(err)
		AddLogger(r.Err, zapcore.Field{Key: "url", String: r.Url})
		return
	}
	if len(r.Params) > 0 {
		urlRow := url.Values{}
		for key, value := range r.Params {
			urlRow.Add(key, value)
		}
		uri.RawQuery = urlRow.Encode()
	}

	r.Request, err = http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		r.Err = ErrReqNew.WithErr(err)
		AddLogger(r.Err, zapcore.Field{Key: "uri", String: uri.String()})
		return
	}

	r.Request.Header = r.Header
	r.client.Timeout = time.Duration(time.Second * 10)
	r.Respose, err = r.client.Do(r.Request)
	defer r.client.CloseIdleConnections()
	if err != nil {
		r.Err = ErrHttpGet.WithErr(err)
		AddLogger(r.Err, zapcore.Field{Key: "url", String: r.Url})
	}
}

func (r *UtilsRequest) Post() {
	uri, err := url.Parse(r.Url)
	if err != nil {
		r.Err = ErrReqParse.WithErr(err)
		AddLogger(r.Err, zapcore.Field{Key: "url", String: r.Url})
		return
	}
	if len(r.Params) > 0 {
		urlRow := url.Values{}
		for key, value := range r.Params {
			urlRow.Add(key, value)
		}
		uri.RawQuery = urlRow.Encode()
	}

	var byteList []byte
	byteList, err = json.Marshal(r.Body)
	if err != nil {
		r.Err = ErrJsonMarshal.WithErr(err)
		AddLogger(r.Err, zapcore.Field{Key: "body", Interface: r.Body})
		return
	}
	r.Request, err = http.NewRequest(http.MethodPost, uri.String(), bytes.NewReader(byteList))
	if err != nil {
		r.Err = ErrReqNew.WithErr(err)
		AddLogger(r.Err, zapcore.Field{Key: "uri", String: uri.String()})
		return
	}

	r.Request.Header = r.Header
	r.client.Timeout = time.Duration(time.Second * 10)
	r.Respose, err = r.client.Do(r.Request)
	defer r.client.CloseIdleConnections()
	if err != nil {
		r.Err = ErrHttpPost.WithErr(err)
		AddLogger(r.Err, zapcore.Field{Key: "url", String: r.Url})
	}
}

func (r *UtilsRequest) GetHttps() {
	var (
		// 创建证书池及各类对象
		pool   *x509.CertPool // 我们要把一部分证书存到这个池中
		client *http.Client
		resp   *http.Response
		err    error
		caCrt  []byte          // 根证书
		cliCrt tls.Certificate // 具体的证书加载对象
	)

	// 读取caCertPath
	caCrt, err = ioutil.ReadFile(r.CaCertPath)
	if err != nil {
		r.Err = ErrReadFile.WithErr(err)
		AddLogger(r.Err)
		return
	}

	// NewCertPool
	pool = x509.NewCertPool()

	// 解析一系列PEM编码的证书
	pool.AppendCertsFromPEM(caCrt)

	// 具体的证书加载对象
	cliCrt, err = tls.LoadX509KeyPair(r.CertFile, r.KeyFile)
	if err != nil {
		r.Err = ErrLoadX509KeyPair.WithErr(err)
		return
	}

	// 把上面的准备内容传入 client
	r.client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      pool,
				Certificates: []tls.Certificate{cliCrt},
			},
		},
	}

	// Get 请求
	resp, err = client.Get(r.Url)
	// 延时关闭
	defer client.CloseIdleConnections()
	if err != nil {
		r.Err = ErrHttpPost.WithErr(err)
		AddLogger(r.Err, zapcore.Field{Key: "url", String: r.Url})
		return
	}
	// 延时关闭
	defer resp.Body.Close()
}
