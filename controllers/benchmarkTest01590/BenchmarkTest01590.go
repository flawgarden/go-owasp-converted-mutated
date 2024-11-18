package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest01590 struct{}

func (b *BenchmarkTest01590) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	values := r.Form["BenchmarkTest01590"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := b.doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	_, err := fmt.Fprintf(w, bar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (b *BenchmarkTest01590) doSomething(r *http.Request, param string) string {
	bar := param
	return bar
}

func main() {
	http.Handle("/xss-03/BenchmarkTest01590", &BenchmarkTest01590{})
	http.ListenAndServe(":8080", nil)
}
