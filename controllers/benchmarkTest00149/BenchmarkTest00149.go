package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest00149 struct{}

func (b *BenchmarkTest00149) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("Referer") != "" {
		param = r.Header.Get("Referer")
	}

	param, _ = url.QueryUnescape(param)

	sbxyz19132 := param + "_SafeStuff"
	bar := sbxyz19132

	w.Header().Set("X-XSS-Protection", "0")
	output := fmt.Sprintf("Formatted like: %s and %s.", "a", bar)
	w.Write([]byte(output))
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00149", &BenchmarkTest00149{})
	http.ListenAndServe(":8080", nil)
}
