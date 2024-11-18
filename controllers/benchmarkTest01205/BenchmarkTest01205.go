package controllers

import (
	"net/http"
	"net/url"
	"text/template"
)

type BenchmarkTest01205 struct{}

func (b *BenchmarkTest01205) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	param := ""
	headers := r.Header["BenchmarkTest01205"]
	if len(headers) > 0 {
		param = headers[0]
	}
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	r.Context().Value("session").(http.ResponseWriter).Write([]byte("Item: 'userid' with value: '" + encodeForHTML(bar) + "' saved in session."))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := ""
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func encodeForHTML(s string) string {
	return template.HTMLEscapeString(s)
}
