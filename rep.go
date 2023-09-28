package BuilderHttpClient

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"golang.org/x/net/html"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"log"
	"reflect"

	"github.com/PuerkitoBio/goquery"
	"strings"
)

type ResponseBuilder struct {
	request    *ClientBuilder
	code       int
	status     string
	bodyResult []byte
	cookie     string
}

type ResponseInterfaceBuilder interface {
	Code() int
	Status() string
	Json(v any) error
	Text() string
	TextGbk() string
	Gjson() gjson.Result
	Debug() *ResponseBuilder
	DebugString() string
	Byte() []byte
	Cookie() string
	Html() *goquery.Document
	HtmlGbk() *goquery.Document
}

func (c *ResponseBuilder) Debug() *ResponseBuilder {
	log.Println(c.DebugString())
	return c
}

func (c *ResponseBuilder) DebugString() string {
	var sb strings.Builder
	sb.WriteString("START ====================================\n")
	if c.request == nil {
		sb.WriteString("请求体为空，无法读取")
		sb.WriteString("END \n====================================\n")
		return sb.String()
	}
	sb.WriteString("Method: ")
	sb.WriteString(c.request.r.Method)
	sb.WriteString("\nURL: ")
	sb.WriteString(c.request.r.URL.String())
	sb.WriteString("\nQuery Data: ")
	sb.WriteString(c.request.requestBody)
	sb.WriteString("\n")

	sb.WriteString("HeaderSTART: ")
	for k, v := range c.request.r.Header {
		sb.WriteString(fmt.Sprintf("\t%s: %s\n", k, v))
	}
	sb.WriteString("HeaderEND\n")
	sb.WriteString("Cookie: ")
	sb.WriteString(c.cookie)
	sb.WriteString("\nStatus: ")
	sb.WriteString(c.status)
	sb.WriteString("\nStatusCode: ")
	sb.WriteString(fmt.Sprintf("%v\n", c.code))

	if c.bodyResult == nil {
		sb.WriteString("响应体为空，无法读取")
	} else {
		sb.WriteString("Result: ")
		sb.WriteString(c.Text())
	}
	sb.WriteString("END \n====================================\n")

	return sb.String()
}
func (c *ResponseBuilder) Code() int {
	return c.code
}
func (c *ResponseBuilder) Status() string {
	return c.status
}

func (c *ResponseBuilder) Json(v any) error {
	if v == nil {
		return fmt.Errorf("DecodeJson:传入的对象不能为空")
	} else if c.bodyResult == nil {
		return fmt.Errorf("DecodeJson:响应体为空，无法读取")
	} else {
		valueType := reflect.TypeOf(v)
		if valueType.Kind() != reflect.Ptr {
			return fmt.Errorf("DecodeJson:传入的对象必须是指针类型")
		}
	}
	return json.NewDecoder(strings.NewReader(c.Text())).Decode(v)
}

func (c *ResponseBuilder) Text() string {
	return string(c.bodyResult)
}
func (c *ResponseBuilder) TextGbk() string {
	decoder := simplifiedchinese.GBK.NewDecoder()
	utf8BodyReader := transform.NewReader(strings.NewReader(c.Text()), decoder)
	utf8Body, err := io.ReadAll(utf8BodyReader)
	if err != nil {
		fmt.Println("解码失败:", err)
		return ""
	}
	return string(utf8Body)
}
func (c *ResponseBuilder) Html() *goquery.Document {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(c.Text()))
	if err != nil {
		log.Printf("读取响应体失败: %s", err)
		return nil
	}
	return doc
}
func (c *ResponseBuilder) HtmlGbk() *goquery.Document {
	docs, err := html.Parse(strings.NewReader(c.TextGbk()))
	if err != nil {
		fmt.Println("解析HTML失败:", err)
		return nil
	}
	doc := goquery.NewDocumentFromNode(docs)
	if err != nil {
		fmt.Println("解析HTML失败:", err)
		return nil
	}
	return doc
}

func (c *ResponseBuilder) Byte() []byte {
	return c.bodyResult
}

func (c *ResponseBuilder) Gjson() gjson.Result {
	return gjson.Parse(c.Text())
}
func (c *ResponseBuilder) Cookie() string {
	return c.cookie
}
