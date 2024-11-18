package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest01587 struct{}

func (b *BenchmarkTest01587) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	responseContentType := "text/html;charset=UTF-8"
	w.Header().Set("Content-Type", responseContentType)

	values := r.Form["BenchmarkTest01587"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := b.doSomething(param)
	w.Header().Set("X-XSS-Protection", "0")
	fmt.Fprintf(w, bar)
}

func (b *BenchmarkTest01587) doSomething(param string) string {
	num := 106
	bar := ""
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}
	return bar
}

func main() {
	http.Handle("/xss-03/BenchmarkTest01587", &BenchmarkTest01587{})
	http.ListenAndServe(":8080", nil)
}
