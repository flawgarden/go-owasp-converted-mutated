package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest02051 struct{}

func (b *BenchmarkTest02051) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	headers := r.Header["Referer"]
	if len(headers) > 0 {
		param = headers[0] // just grab first element
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{bar, "b"}
	_, _ = fmt.Fprintf(w, "Formatted like: %1$s and %2$s.", obj)
}

func doSomething(r *http.Request, param string) string {
	var bar string

	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

	return bar
}
