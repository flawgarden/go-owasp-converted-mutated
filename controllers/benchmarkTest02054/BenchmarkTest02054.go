package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest02054 struct{}

func (bt *BenchmarkTest02054) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	headers := r.Header["Referer"]

	if len(headers) > 0 {
		param = headers[0]
	}

	decodedParam, _ := url.QueryUnescape(param)
	bar := doSomething(r, decodedParam)

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", "b"}
	fmt.Fprintf(w, bar, obj...)
}

func doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		bar = param
	}

	return bar
}
