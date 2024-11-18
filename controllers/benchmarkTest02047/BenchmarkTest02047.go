package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest02047 struct{}

func (bt *BenchmarkTest02047) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	var param string

	if headers := r.Header["Referer"]; len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func doSomething(r *http.Request, param string) string {
	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}

func main() {
	http.Handle("/xss-03/BenchmarkTest02047", &BenchmarkTest02047{})
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
