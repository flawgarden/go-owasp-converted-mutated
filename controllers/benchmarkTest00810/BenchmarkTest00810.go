package controllers

import (
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest00810 struct{}

func (b *BenchmarkTest00810) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.RawQuery
	paramVal := "BenchmarkTest00810="
	paramLoc := strings.Index(queryString, paramVal)

	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest00810' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramVal):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramVal) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := param + "_SafeStuff"

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func main() {
	http.Handle("/xss-01/BenchmarkTest00810", &BenchmarkTest00810{})
	http.ListenAndServe(":8080", nil)
}