package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest01507 struct{}

func (b *BenchmarkTest01507) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01507")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", "b"}
	fmt.Fprintf(w, bar, obj...)
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	var bar string

	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	return bar
}
