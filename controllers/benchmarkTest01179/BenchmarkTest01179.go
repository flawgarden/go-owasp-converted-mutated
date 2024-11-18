package controllers

import (
	"net/http"
	"net/url"
)

type BenchmarkTest01179 struct{}

func (b *BenchmarkTest01179) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if referer := r.Header.Get("Referer"); referer != "" {
		param = referer
	}

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		bar = param
	}
	return bar
}

func main() {
	http.Handle("/xss-02/BenchmarkTest01179", &BenchmarkTest01179{})
	http.ListenAndServe(":8080", nil)
}
