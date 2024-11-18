package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest00157 struct{}

func (b *BenchmarkTest00157) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("Referer") != "" {
		param = r.Header.Get("Referer")
	}

	param, _ = url.QueryUnescape(param)

	bar := ""
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

	w.Header().Set("X-XSS-Protection", "0")
	_, _ = w.Write([]byte(bar))
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00157", &BenchmarkTest00157{})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
