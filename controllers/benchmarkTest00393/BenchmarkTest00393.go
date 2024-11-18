package controllers

import (
	"net/http"
)

type BenchmarkTest00393 struct{}

func (b *BenchmarkTest00393) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00393")
	if param == "" {
		param = ""
	}

	var bar string
	num := 106

	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	w.Header().Set("X-XSS-Protection", "0")
	_, _ = w.Write([]byte(bar))
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00393", &BenchmarkTest00393{})
	http.ListenAndServe(":8080", nil)
}
