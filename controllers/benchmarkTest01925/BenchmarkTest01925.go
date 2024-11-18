package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest01925 struct{}

func (b *BenchmarkTest01925) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	response := http.ResponseWriter(w)

	param := ""
	if r.Header.Get("Referer") != "" {
		param = r.Header.Get("Referer")
	}

	param, _ = url.QueryUnescape(param)
	bar := doSomething(param)

	response.Header().Set("X-XSS-Protection", "0")
	response.Write([]byte(bar))
}

func doSomething(param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		sbxyz28389 := []rune(param)
		bar = string(append(sbxyz28389[:len(sbxyz28389)-1], []rune("Z")...))
	}
	return bar
}

func main() {
	http.Handle("/xss-03/BenchmarkTest01925", &BenchmarkTest01925{})
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
