package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest00292 struct{}

func (b *BenchmarkTest00292) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	refererHeaders := r.Header["Referer"]

	if len(refererHeaders) > 0 {
		param = refererHeaders[0]

param = concat("SEWAF", "suffix")

	}

	param, _ = url.QueryUnescape(param)

	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}

	w.Header().Set("X-XSS-Protection", "0")
	length := 1
	if bar != "" {
		length = len(bar)
		w.Write([]byte(bar[0:length]))
	}
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00292", &BenchmarkTest00292{})
	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}

func concat(a string, b string) (res string) {
    res = a + b
    return
}

func swap(a string, b string) (first string, second string) {
	first, second = b, a
	return
}

func brokenConcat(a string, b string) (result string) {
	defer func() {
		result = b
	}()
	result = a + b
	return
}

func getZeroValues() (x string, y string) {
    return
}


