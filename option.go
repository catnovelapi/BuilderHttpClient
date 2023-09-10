package buildHttps

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

type Option interface {
	apply(*ClientBuilder)
}

type OptionFunc func(*ClientBuilder)

func (optionFunc OptionFunc) apply(c *ClientBuilder) {
	optionFunc(c)
}

func Debug() Option {
	return OptionFunc(func(c *ClientBuilder) {
		file, err := os.OpenFile("http_builder.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Println(err)
		} else {
			log.SetOutput(file)
			log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
			log.SetPrefix("[http_builder]")
		}
	})
}
func Method(method string) Option {
	return OptionFunc(func(c *ClientBuilder) {
		c.Method = method
	})
}
func ApiPath(URL string) Option {
	return OptionFunc(func(c *ClientBuilder) {
		u, err := url.Parse(URL)
		if err != nil {
			log.Printf("url 解析失败: %s", err)
		}
		if strings.HasSuffix(u.Host, ":") {
			u.Host = strings.TrimSuffix(u.Host, ":")
		}
		c.URL = u
		c.Host = u.Host
	})
}
func Header(header map[string]any) Option {
	return OptionFunc(func(c *ClientBuilder) {
		if header != nil {
			for k, v := range header {
				c.Header.Set(k, fmt.Sprintf("%v", v))
			}
		}
	})

}
func Body(dataBody any) Option {
	return OptionFunc(func(c *ClientBuilder) {
		c.dataBody = dataBody
	})
}
