package controllers

import (
	"net/http"
)

type BenchmarkTest02482 struct{}

func (b *BenchmarkTest02482) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := r.Form["BenchmarkTest02482"]
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func doSomething(param string) string {
	bar := ""
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	return bar
}

func main() {
	http.Handle("/xss-04/BenchmarkTest02482", &BenchmarkTest02482{})
	http.ListenAndServe(":8080", nil)
}
