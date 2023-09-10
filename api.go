package buildHttps

import (
	"net/http"
)

func httpDefault(options []Option, optionsDefault ...Option) ResponseInterfaceBuilder {
	var optionArray []Option
	for _, opt := range optionsDefault {
		optionArray = append(optionArray, opt)
	}
	if options != nil {
		for _, opt := range options {
			optionArray = append(optionArray, opt)
		}
	}
	return Df(optionArray...).Build()
}

func Get(url string, options ...Option) ResponseInterfaceBuilder {
	return httpDefault(options, ApiPath(url))
}

func Post(url string, options ...Option) ResponseInterfaceBuilder {
	return httpDefault(options, Method(http.MethodPost), ApiPath(url))
}

func Put(url string, options ...Option) ResponseInterfaceBuilder {
	return httpDefault(options, Method(http.MethodPut), ApiPath(url))
}

func Delete(url string, options ...Option) ResponseInterfaceBuilder {
	return httpDefault(options, Method(http.MethodDelete), ApiPath(url))
}

func Patch(url string, options ...Option) ResponseInterfaceBuilder {
	return httpDefault(options, Method(http.MethodPatch), ApiPath(url))
}
