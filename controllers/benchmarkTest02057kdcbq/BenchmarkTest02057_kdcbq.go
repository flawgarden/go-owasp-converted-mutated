package controllers

import (
	"net/http"
	"net/url"
)

type BenchmarkTest02057 struct{}

func (b *BenchmarkTest02057) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := ""

	if referrer := r.Header.Get("Referer"); referrer != "" {
		param = referrer
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")

var i123 interface{} = &EmbeddedStruct{Field1: bar}
if ptr, ok := i123.(*EmbeddedStruct); ok {
     bar = ptr.Field1
} else {
    bar = "KJjrm"
}

	if bar != "" {
		w.Write([]byte(bar))
	}
}

func doSomething(param string) string {
	bar := ""
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	return bar
}
