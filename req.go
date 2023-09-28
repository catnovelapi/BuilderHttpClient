package BuilderHttpClient

import (
	"fmt"
	"io"
	"net/url"
	"sync"
	"time"

	"net/http"
)

type ClientBuilder struct {
	sync.Mutex
	r           *http.Request
	timeout     time.Duration
	proxyPath   *url.URL
	dataBody    any
	requestBody string
}

type RequestInterfaceBuilder interface {
	NewRequestClient() (*ResponseBuilder, error)
}

func (b *ClientBuilder) NewRequestClient() (*ResponseBuilder, error) {
	defaultClient := &http.Client{}
	if b.proxyPath != nil {
		defaultClient.Transport = &http.Transport{Proxy: http.ProxyURL(b.proxyPath)}
	}
	if b.timeout != 0 {
		defaultClient.Timeout = b.timeout
	}

	do, err := defaultClient.Do(b.r)

	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %s", err)
	} else if do == nil {
		return nil, fmt.Errorf("请求失败,响应体为空")
	} else if do.Body == nil {
		return nil, fmt.Errorf("响应体Body为空")
	} else {
		defer do.Body.Close()
	}

	rep := &ResponseBuilder{
		request: b,
		code:    do.StatusCode,
		status:  do.Status,
	}
	body, ok := io.ReadAll(do.Body)
	if ok != nil {
		return nil, fmt.Errorf("读取响应体失败: %s", ok)
	}
	rep.bodyResult = body
	for _, cookie := range do.Cookies() {
		rep.cookie += cookie.Name + "=" + cookie.Value + ";"
	}
	return rep, nil
}
