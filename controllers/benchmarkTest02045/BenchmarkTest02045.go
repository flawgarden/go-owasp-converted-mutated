package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest02045 struct{}

func (b *BenchmarkTest02045) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	var param string

	referer := r.Header.Get("Referer")
	if referer != "" {
		param = referer
	}

	param, _ = url.QueryUnescape(param)
	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", bar}
	output := "<!DOCTYPE html>\n<html>\n<body>\n<p>" + formatOutput(obj) + "\n</p>\n</body>\n</html>"
	w.Write([]byte(output))
}

func doSomething(param string) string {
	var bar string
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func formatOutput(obj []interface{}) string {
	return fmt.Sprintf("Formatted like: %s and %s.", obj[0], obj[1])
}
