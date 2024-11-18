package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest00146 struct{}

func (b *BenchmarkTest00146) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("Referer")
	if param != "" {
		// URL Decode the header value
		param, _ = url.QueryUnescape(param)
	}

	bar := param
	if len(param) > 1 {
		sbxyz67327 := []rune(param)
		bar = string(append(sbxyz67327[:len(param)-1], 'Z'))
	}

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", "b"}
	_, _ = w.Write([]byte(fmt.Sprintf(bar, obj...)))
}
