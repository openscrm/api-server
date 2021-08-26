package workwx

import (
	"github.com/go-resty/resty/v2"
	"net/http"
)

// DefaultQYAPIHost 默认企业微信 API Host
const DefaultQYAPIHost = "https://qyapi.weixin.qq.com"

type options struct {
	WxAPIHost string
	HTTP      *http.Client
	restyCli  *resty.Client
}

// CtorOption 客户端对象构造参数
type CtorOption interface {
	applyTo(*options)
}

// impl Default for options
func defaultOptions() (opt options) {
	opt = options{
		WxAPIHost: DefaultQYAPIHost,
		HTTP:      &http.Client{},
	}
	opt.restyCli = resty.NewWithClient(opt.HTTP)
	return
}

type withQYAPIHost struct {
	x string
}

// WithQYAPIHost 覆盖默认企业微信 API 域名
func WithQYAPIHost(host string) CtorOption {
	return &withQYAPIHost{x: host}
}

var _ CtorOption = (*withQYAPIHost)(nil)

func (x *withQYAPIHost) applyTo(y *options) {
	y.WxAPIHost = x.x
}

type withHTTPClient struct {
	x *http.Client
}

// WithHTTPClient 使用给定的 http.Client 作为 HTTP 客户端
func WithHTTPClient(client *http.Client) CtorOption {
	return &withHTTPClient{x: client}
}

var _ CtorOption = (*withHTTPClient)(nil)

func (x *withHTTPClient) applyTo(y *options) {
	y.HTTP = x.x
	y.restyCli = resty.NewWithClient(x.x)
}
