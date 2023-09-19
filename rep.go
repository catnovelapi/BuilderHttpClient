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

	"github.com/PuerkitoBio/goquery"
	"strings"
)

type ResponseBuilder struct {
	request    *ClientBuilder
	code       int
	status     string
	body       io.Reader
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
	sb.WriteString("Method: ")
	sb.WriteString(c.request.Method)
	sb.WriteString("\nURL: ")
	sb.WriteString(c.request.URL.String())
	sb.WriteString("\nQuery Data: ")
	sb.WriteString(c.request.requestBody)
	sb.WriteString("\n")

	sb.WriteString("HeaderSTART: ")
	for k, v := range c.request.Header {
		sb.WriteString(fmt.Sprintf("\t%s: %s\n", k, v))
	}
	sb.WriteString("HeaderEND\n")
	sb.WriteString("Cookie: ")
	sb.WriteString(c.cookie)
	sb.WriteString("\nStatus: ")
	sb.WriteString(c.status)
	sb.WriteString("\nStatusCode: ")
	sb.WriteString(fmt.Sprintf("%v\n", c.code))

	if c.body == nil {
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
	return string(c.Byte())
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
	doc, err := goquery.NewDocumentFromReader(c.body)
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
	if c.body == nil {
		log.Printf("响应体为空,无法读取")
		return nil
	}
	if c.bodyResult != nil {
		return c.bodyResult
	}
	b, err := io.ReadAll(c.body)
	if err != nil {
		log.Printf("读取响应体失败: %s", err)
		return nil
	}
	c.bodyResult = b
	return b
}

func (c *ResponseBuilder) Gjson() gjson.Result {
	return gjson.ParseBytes(c.Byte())
}
func (c *ResponseBuilder) Cookie() string {
	return c.cookie
}
