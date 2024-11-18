package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest00477 struct{}

func (b *BenchmarkTest00477) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00477")
	bar := fmt.Sprintf("%s_SafeStuff", param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00477", &BenchmarkTest00477{})
	http.ListenAndServe(":8080", nil)
}
