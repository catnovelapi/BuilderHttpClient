package BuilderHttpClient

import (
	"net/http"
)

func BuilderDefault(url string, options ...Option) *ClientBuilder {
	options = append(options, ApiPath(url))
	req := &ClientBuilder{}
	req.Header = make(http.Header)
	req.Method = http.MethodGet
	req.Proto = "HTTP/1.1"
	req.ProtoMajor = 1
	req.ProtoMinor = 1
	Timeout(10).apply(req)
	for _, opt := range options {
		opt.apply(req)
	}
	return req
}

func Get(url string, options ...Option) ResponseInterfaceBuilder {
	return BuilderDefault(url, options...).NewBuild()
}

func Post(url string, options ...Option) ResponseInterfaceBuilder {
	options = append(options, Method(http.MethodPost))
	return BuilderDefault(url, options...).NewBuild()
}

func Put(url string, options ...Option) ResponseInterfaceBuilder {
	options = append(options, Method(http.MethodPut))
	return BuilderDefault(url, options...).NewBuild()
}

func Delete(url string, options ...Option) ResponseInterfaceBuilder {
	options = append(options, Method(http.MethodDelete))
	return BuilderDefault(url, options...).NewBuild()
}

func Patch(url string, options ...Option) ResponseInterfaceBuilder {
	options = append(options, Method(http.MethodPatch))
	return BuilderDefault(url, options...).NewBuild()
}
