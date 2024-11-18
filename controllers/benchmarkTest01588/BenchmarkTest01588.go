package controllers

import (
	"net/http"
)

type BenchmarkTest01588 struct{}

func (bt *BenchmarkTest01588) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	values := r.URL.Query()["BenchmarkTest01588"]
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	bar := bt.doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	_, _ = w.Write([]byte(bar))
}

func (bt *BenchmarkTest01588) doSomething(r *http.Request, param string) string {
	var bar string
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
