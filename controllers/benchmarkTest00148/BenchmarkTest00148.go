package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest00148Controller struct {
	http.Handler
}

func (c *BenchmarkTest00148Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("Referer") != "" {
		param = r.Header.Get("Referer")
	}

	param, _ = url.QueryUnescape(param)

	var bar string

	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", bar}
	_, _ = fmt.Fprintf(w, "Formatted like: %1$s and %2$s.", obj)
}
