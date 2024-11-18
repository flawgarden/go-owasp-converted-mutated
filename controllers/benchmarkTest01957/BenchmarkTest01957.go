package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest01957 struct{}

func (b *BenchmarkTest01957) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("BenchmarkTest01957") != "" {
		param = r.Header.Get("BenchmarkTest01957")
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(r, param)

	r.AddCookie(&http.Cookie{Name: bar, Value: "10340"})

	fmt.Fprintf(w, "Item: '%s' with value: 10340 saved in session.", bar)
}

func doSomething(r *http.Request, param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[1]

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
