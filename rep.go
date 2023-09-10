package BuilderHttpClient

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"io"
	"log"
)

type ResponseBuilder struct {
	request    *ClientBuilder
	code       int
	status     string
	body       io.Reader
	bodyString string
	cookie     string
}
type ResponseInterfaceBuilder interface {
	Code() int
	Status() string
	Json(v any) error
	Text() string
	Gjson() gjson.Result
	Debug() *ResponseBuilder
}

func (c *ResponseBuilder) Debug() *ResponseBuilder {
	log.Println("START ====================================")
	log.Printf("Method: %s\n", c.request.Method)
	log.Printf("URL: %s\n", c.request.URL.String())
	log.Printf("Query Data: %s\n", c.request.requestBody)
	for k, v := range c.request.Header {
		log.Printf("Header :  %s: %s\n", k, v)
	}
	log.Printf("Cookie : %s\n", c.cookie)
	log.Printf("Status: %s\n", c.status)
	log.Printf("StatusCode: %v\n", c.code)
	if c.body == nil {
		log.Println("响应体为空,无法读取")
	} else {
		log.Printf("Result: %s\n", c.Text())
	}
	log.Println("==================================== END")
	return c
}
func (c *ResponseBuilder) Code() int {
	return c.code
}
func (c *ResponseBuilder) Status() string {
	return c.status
}

func (c *ResponseBuilder) Json(v any) error {
	if c.body == nil {
		log.Printf("DecodeJson:响应体为空")
		return nil
	}
	if v == nil {
		log.Printf("DecodeJson:传入的对象不能为空")
		return nil
	}
	return json.NewDecoder(c.body).Decode(v)
}

func (c *ResponseBuilder) Text() string {
	if c.body == nil {
		log.Printf("响应体为空,无法读取")
		return ""
	}
	if c.bodyString != "" {
		return c.bodyString
	}
	b, err := io.ReadAll(c.body)
	if err != nil {
		log.Printf("读取响应体失败: %s", err)
		return ""
	}
	c.bodyString = string(b)
	return c.bodyString
}

func (c *ResponseBuilder) Gjson() gjson.Result {
	return gjson.Parse(c.Text())
}
