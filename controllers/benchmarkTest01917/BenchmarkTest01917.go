package controllers

import (
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest01917 struct{}

func (b *BenchmarkTest01917) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.Header.Get("Referer")
	if param != "" {
		param = decode(param)
	}

	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func doSomething(param string) string {
	bar := encodeForHTML(param)
	return bar
}

func encodeForHTML(input string) string {
	return strings.Replace(strings.Replace(input, "&", "&amp;", -1), "<", "&lt;", -1)
}

func decode(input string) string {
	decoded, _ := url.QueryUnescape(input)
	return decoded
}

func main() {
	http.Handle("/xss-03/BenchmarkTest01917", &BenchmarkTest01917{})
	http.ListenAndServe(":8080", nil)
}
