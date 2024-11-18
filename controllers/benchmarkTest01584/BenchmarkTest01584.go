package controllers

import (
	"net/http"
	"strings"
)

type BenchmarkTest01584 struct{}

func (b *BenchmarkTest01584) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0")
	r.ParseForm()
	values := r.Form["BenchmarkTest01584"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}
	bar := new(Test).doSomething(r, param)
	obj := []interface{}{"a", "b"}
	w.Write([]byte(formatString(bar, obj)))
}

func formatString(bar string, obj []interface{}) string {
	return strings.Replace(bar, "%s", obj[0].(string), 1)
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := ""
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}
