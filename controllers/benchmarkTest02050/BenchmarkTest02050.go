package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest02050 struct{}

func (bt *BenchmarkTest02050) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	headers := r.Header["Referer"]
	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)
	bar := doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", bar}
	fmt.Fprintf(w, "Formatted like: %s and %s.", obj[0], obj[1])
}

func doSomething(r *http.Request, param string) string {
	bar := ""
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}

func main() {
	http.Handle("/xss-03/BenchmarkTest02050", &BenchmarkTest02050{})
	http.ListenAndServe(":8080", nil)
}
