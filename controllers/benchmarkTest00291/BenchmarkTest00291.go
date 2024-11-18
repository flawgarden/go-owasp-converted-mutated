package controllers

import (
	"net/http"
	"net/url"
)

type BenchmarkTest00291 struct{}

func (b *BenchmarkTest00291) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	headers := r.Header["Referer"]

	if len(headers) > 0 {
		param = headers[0] // just grab first element
	}

	param, _ = url.QueryUnescape(param)

	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}
