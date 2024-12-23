package controllers

import (
	"net/http"
	"strings"
)

type BenchmarkTest02228 struct{}

func (b *BenchmarkTest02228) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest02228")

nested7231 := NewNestedFields4(param)
param = nested7231.nested1.nested1.nested1.nested1.value

	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func doSomething(param string) string {
	var sb strings.Builder
	sb.WriteString(param)
	bar := sb.String() + "_SafeStuff"
	return bar
}

func main() {
	http.Handle("/xss-04/BenchmarkTest02228", &BenchmarkTest02228{})
	http.ListenAndServe(":8080", nil)
}
