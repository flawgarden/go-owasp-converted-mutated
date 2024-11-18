package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest01918 struct{}

func (b *BenchmarkTest01918) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.Header.Get("Referer")

	if param != "" {
		decodedParam, _ := url.QueryUnescape(param)
		param = decodedParam
	}

	bar := b.doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", bar}
	_, _ = w.Write([]byte(fmt.Sprintf("Formatted like: %1$s and %2$s.", obj...)))
}

func (b *BenchmarkTest01918) doSomething(r *http.Request, param string) string {
	var bar string
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
