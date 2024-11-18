package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest01915 struct{}

func (b *BenchmarkTest01915) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if referer := r.Header.Get("Referer"); referer != "" {
		param = referer
	}

	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	_, _ = fmt.Fprintf(w, "Formatted like: %s and %s.", "a", bar)
}

func (b *BenchmarkTest01915) doSomething(r *http.Request, param string) string {
	bar := ""
	num := 106

	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}
