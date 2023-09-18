package BuilderHttpClient

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type ClientBuilder struct {
	http.Request
	dataBody    any
	requestBody string
}

type RequestInterfaceBuilder interface {
	Build() ResponseInterfaceBuilder
}

func (b *ClientBuilder) NewCookie(cookie map[string]string) *ClientBuilder {
	for k, v := range cookie {
		s := fmt.Sprintf("%s=%s", k, v)
		if c := b.Header.Get("Cookie"); c != "" {
			b.Header.Set("Cookie", c+"; "+s)
		} else {
			b.Header.Set("Cookie", s)
		}
	}
	return b
}

func (b *ClientBuilder) Build() *ResponseBuilder {
	if b.Method == http.MethodGet {
		b.requestBody = encodeDataFormValue(b.dataBody)
		b.URL.RawQuery = b.requestBody
		b.Header.Del("Content-Type")
	} else {
		if b.Header.Get("Content-Type") == "" {
			b.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
		switch b.Header.Get("Content-Type") {
		case "application/json":
			b.requestBody = encodeJsonDataValue(b.dataBody)
		case "multipart/form-data":
		default:
			b.requestBody = encodeDataFormValue(b.dataBody)
		}
		b.Body = io.NopCloser(strings.NewReader(b.requestBody))
	}

	sendClient, err := http.DefaultClient.Do(&b.Request)
	if err != nil {
		log.Printf("发送请求失败: %s", err)
		return &ResponseBuilder{}
	}
	responseClient := &ResponseBuilder{
		request: b,
		code:    sendClient.StatusCode,
		status:  sendClient.Status,
		body:    sendClient.Body,
	}
	for _, cookie := range sendClient.Cookies() {
		responseClient.cookie += cookie.Name + "=" + cookie.Value + ";"
	}
	return responseClient
}

func (b *ClientBuilder) NewBuild() ResponseInterfaceBuilder {
	return b.Build()
}
