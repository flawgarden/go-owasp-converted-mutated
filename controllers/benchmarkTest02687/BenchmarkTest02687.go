package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest02687 struct{}

func (b *BenchmarkTest02687) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02687")

	bar := doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprintf(w, bar, "a", "b")
}

func doSomething(r *http.Request, param string) string {
	bar := ""
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func main() {
	http.Handle("/xss-05/BenchmarkTest02687", &BenchmarkTest02687{})
	http.ListenAndServe(":8080", nil)
}
