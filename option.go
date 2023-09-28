package BuilderHttpClient

import (
	"fmt"
	"log"
	"net/url"
	"time"
)

type Option interface {
	apply(*ClientBuilder)
}

type OptionFunc func(*ClientBuilder)

func (optionFunc OptionFunc) apply(c *ClientBuilder) {
	optionFunc(c)
}

func Cookie(cookie map[string]string) Option {
	return OptionFunc(func(b *ClientBuilder) {
		for k, v := range cookie {
			s := fmt.Sprintf("%s=%s", k, v)
			if c := b.r.Header.Get("Cookie"); c != "" {
				b.r.Header.Set("Cookie", c+"; "+s)
			} else {
				b.r.Header.Set("Cookie", s)
			}
		}
	})
}

func Timeout(timeout int) Option {
	return OptionFunc(func(c *ClientBuilder) {
		c.timeout = time.Duration(timeout) * time.Second
	})
}

func Proxy(proxy string) Option {
	return OptionFunc(func(c *ClientBuilder) {
		c.Lock()
		defer c.Unlock()
		var err error
		c.proxyPath, err = url.Parse(proxy)
		if err != nil {
			log.Printf("proxy 解析失败: %s", err)
			c.proxyPath = nil
		}
	})
}
func Method(method string) Option {
	return OptionFunc(func(c *ClientBuilder) {
		c.r.Method = method
	})
}
func Header(header map[string]any) Option {
	return OptionFunc(func(c *ClientBuilder) {
		if header != nil {
			for k, v := range header {
				c.r.Header.Set(k, fmt.Sprintf("%v", v))
			}
		}
	})

}
func Body(dataBody any) Option {
	return OptionFunc(func(c *ClientBuilder) {
		c.dataBody = dataBody
	})
}
