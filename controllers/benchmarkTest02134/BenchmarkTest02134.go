package controllers

import (
	"net/http"
)

type BenchmarkTest02134 struct{}

func (b *BenchmarkTest02134) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0")
	param := r.URL.Query().Get("BenchmarkTest02134")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte(bar))
}

func doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}

func main() {
	http.Handle("/xss-04/BenchmarkTest02134", &BenchmarkTest02134{})
	http.ListenAndServe(":8080", nil)
}
