package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest01082 struct{}

func (b *BenchmarkTest01082) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("BenchmarkTest01082") != "" {
		param = r.Header.Get("BenchmarkTest01082")
	}

	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(param)

	r.Header.Set(bar, "10340")

	fmt.Fprintf(w, "Item: '%s' with value: '10340' saved in session.", bar)
}

func (b *BenchmarkTest01082) doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}
