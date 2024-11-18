package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest01914 struct{}

func (b *BenchmarkTest01914) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	var param string
	if referer := r.Header.Get("Referer"); referer != "" {
		param = referer
	}

	param, _ = url.QueryUnescape(param)
	bar := b.doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", bar}
	fmt.Fprintf(w, "<!DOCTYPE html>\n<html>\n<body>\n<p>")
	fmt.Fprintf(w, "Formatted like: %s and %s.", obj[0], obj[1])
	fmt.Fprintf(w, "\n</p>\n</body>\n</html>")
}

func (b *BenchmarkTest01914) doSomething(r *http.Request, param string) string {
	var bar string
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
