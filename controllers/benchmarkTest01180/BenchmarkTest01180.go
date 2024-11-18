package controllers

import (
	"net/http"
	"net/url"
)

type BenchmarkTest01180 struct{}

func (b *BenchmarkTest01180) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	var param string
	headers := r.Header["Referer"]
	if len(headers) > 0 {
		param = headers[0]
	}
	param, _ = url.QueryUnescape(param)
	bar := b.doSomething(r, param)
	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func (b *BenchmarkTest01180) doSomething(r *http.Request, param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe
	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}
	return bar
}
