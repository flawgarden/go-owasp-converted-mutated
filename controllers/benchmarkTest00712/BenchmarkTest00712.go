package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest00712 struct {
}

func (b *BenchmarkTest00712) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	values := r.URL.Query()["BenchmarkTest00712"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	var bar string
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", bar}
	response := fmt.Sprintf("<!DOCTYPE html>\n<html>\n<body>\n<p>Formatted like: %s and %s.</p>\n</body>\n</html>", obj[0], obj[1])
	w.Write([]byte(response))
}
