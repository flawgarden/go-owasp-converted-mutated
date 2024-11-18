package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest02679 struct{}

func (bt *BenchmarkTest02679) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest02679")
	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", bar}
	_, _ = fmt.Fprintf(w, "<!DOCTYPE html>\n<html>\n<body>\n<p>Formatted like: %s and %s.\n</p>\n</body>\n</html>", obj[0], obj[1])
}

func doSomething(param string) string {
	bar := ""
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}
	return bar
}

func main() {
	http.Handle("/xss-05/BenchmarkTest02679", &BenchmarkTest02679{})
	http.ListenAndServe(":8080", nil)
}
