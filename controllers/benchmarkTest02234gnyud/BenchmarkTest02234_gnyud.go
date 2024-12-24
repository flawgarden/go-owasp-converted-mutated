package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest02234 struct{}

func (b *BenchmarkTest02234) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var param string
	if values, ok := r.Form["BenchmarkTest02234"]; ok && len(values) > 0 {
		param = values[0]
	}

	bar := doSomething(param)

value := -1
switch {
case value < 0:
    bar = bar + "_suffix"
case value == 0:
    bar = "fixed_string"
    fallthrough
case value > 0:
    bar += " Positive"
default:
    bar = "Unknown"
}

	w.Header().Set("X-XSS-Protection", "0")
	_, _ = fmt.Fprintf(w, bar)
}

func doSomething(param string) string {
	sbxyz86132 := param + "_SafeStuff"
	return sbxyz86132
}