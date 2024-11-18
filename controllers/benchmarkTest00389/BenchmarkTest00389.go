package controllers

import (
	"fmt"
	"net/http"
	"strings"
)

type BenchmarkTest00389 struct{}

func (b *BenchmarkTest00389) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest00389")
	if param == "" {
		param = ""
	}

	bar := escapeHTML(param)

	w.Header().Set("X-XSS-Protection", "0")
	fmt.Fprintln(w, bar)
}

func escapeHTML(input string) string {
	return strings.ReplaceAll(strings.ReplaceAll(input, "&", "&amp;"), "<", "&lt;")
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00389", &BenchmarkTest00389{})
	http.ListenAndServe(":8080", nil)
}
