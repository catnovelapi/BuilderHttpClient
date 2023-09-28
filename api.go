package BuilderHttpClient

import (
	"fmt"
	"golang.org/x/net/context"
	"io"
	"log"
	"net/http"
	"strings"
)

func BuilderDefault(method, url string, options ...Option) (*ClientBuilder, error) {
	request, err := http.NewRequestWithContext(context.Background(), method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败 NewRequestWithContext: %s", err)
	}
	req := &ClientBuilder{r: request}
	for _, opt := range options {
		opt.apply(req)
	}

	if req.r.Method == http.MethodGet {
		req.requestBody = encodeDataFormValue(req.dataBody)
		req.r.URL.RawQuery = req.requestBody
		req.r.Header.Del("Content-Type")
		return req, nil
	}

	contentType := req.r.Header.Get("Content-Type")
	if contentType == "" {
		req.r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	switch contentType {
	case "application/json":
		req.requestBody = encodeJsonDataValue(req.dataBody)
	case "multipart/form-data":
		// Handle multipart/form-data if needed
	default:
		req.requestBody = encodeDataFormValue(req.dataBody)
	}
	req.r.Body = io.NopCloser(strings.NewReader(req.requestBody))

	return req, nil
}

func Df(method, url string, options ...Option) (*ResponseBuilder, error) {
	res, err := BuilderDefault(method, url, options...)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败 BuilderDefault: %s", err)
	} else {
		return res.NewRequestClient()
	}
}
func Get(url string, options ...Option) ResponseInterfaceBuilder {
	if client, err := Df(http.MethodGet, url, options...); err != nil {
		log.Println(err)
		return nil
	} else {
		return client
	}
}

func Post(url string, options ...Option) ResponseInterfaceBuilder {
	if client, err := Df(http.MethodPost, url, options...); err != nil {
		log.Println(err)
		return nil
	} else {
		return client
	}
}

func Put(url string, options ...Option) ResponseInterfaceBuilder {
	client, err := Df(http.MethodPut, url, options...)
	if err != nil {
		log.Println(err)
		return nil
	}
	return client
}

func Delete(url string, options ...Option) ResponseInterfaceBuilder {
	client, err := Df(http.MethodDelete, url, options...)
	if err != nil {
		log.Println(err)
		return nil
	}
	return client
}

func Patch(url string, options ...Option) ResponseInterfaceBuilder {
	client, err := Df(http.MethodPatch, url, options...)
	if err != nil {
		log.Println(err)
		return nil
	}
	return client
}
