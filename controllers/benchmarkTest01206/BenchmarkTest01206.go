package controllers

import (
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest01206 struct{}

func (b *BenchmarkTest01206) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := ""
	if headers := r.Header["BenchmarkTest01206"]; len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	r.AddCookie(&http.Cookie{Name: "userid", Value: bar})

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte("Item: 'userid' with value: '" + htmlEscape(bar) + "' saved in session."))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		sbxyz14623 := strings.Builder{}
		sbxyz14623.WriteString(param)
		bar = sbxyz14623.String()[:len(param)-1] + "Z"
	}
	return bar
}

func htmlEscape(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}
