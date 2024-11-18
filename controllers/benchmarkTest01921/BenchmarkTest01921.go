package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest01921 struct{}

func (bt *BenchmarkTest01921) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("Referer") != "" {
		param = r.Header.Get("Referer")
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{bar, "b"}
	fmt.Fprintf(w, "Formatted like: %1$s and %2$s.", obj)
}

func doSomething(r *http.Request, param string) string {
	var bar string
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
