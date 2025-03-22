package gosthttp

import (
	"net/http"
	"net/http/cookiejar"
)

func NewHttpClientFromHandler(handler http.Handler) http.Client {
	cookiejar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	return http.Client{
		Transport: TestRoundTripper{Handler: handler},
		Jar:       cookiejar,
	}
}

func NewHttpClient() http.Client {
	cookiejar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	return http.Client{
		Jar: cookiejar,
	}
}
