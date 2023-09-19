package BuilderHttpClient

import (
	"io"

	"log"
	"net/http"
	"strings"
)

type ClientBuilder struct {
	http.Request
	http.Client
	dataBody    any
	requestBody string
}

type RequestInterfaceBuilder interface {
	Build() ResponseInterfaceBuilder
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
	sendClient, err := b.Client.Do(&b.Request)
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
