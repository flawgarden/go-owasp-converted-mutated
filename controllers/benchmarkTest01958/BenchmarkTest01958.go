package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest01958 struct{}

func (b *BenchmarkTest01958) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("BenchmarkTest01958") != "" {
		param = r.Header.Get("BenchmarkTest01958")
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(r, param)

	r.Header.Set("userid", bar)

	fmt.Fprintf(w, "Item: 'userid' with value: '%s' saved in session.", htmlEscape(bar))
}

func doSomething(r *http.Request, param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		sbxyz15757 := strings.Builder{}
		sbxyz15757.WriteString(param)
		bar = sbxyz15757.String()[:len(param)-1] + "Z"
	}
	return bar
}

func htmlEscape(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}
