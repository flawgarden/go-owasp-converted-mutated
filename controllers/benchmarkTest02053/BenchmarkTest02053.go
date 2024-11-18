package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest02053 struct{}

func (bt *BenchmarkTest02053) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := ""

	headers := r.Header["Referer"]
	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	output := fmt.Sprintf("Formatted like: %s and %s.", bar, "b")
	w.Write([]byte(output))
}

func doSomething(r *http.Request, param string) string {
	var bar string
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func main() {
	http.Handle("/xss-03/BenchmarkTest02053", &BenchmarkTest02053{})
	http.ListenAndServe(":8080", nil)
}
