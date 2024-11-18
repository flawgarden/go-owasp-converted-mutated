package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest00711 struct{}

func (b *BenchmarkTest00711) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := r.URL.Query()["BenchmarkTest00711"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", bar}
	output := fmt.Sprintf("<!DOCTYPE html>\n<html>\n<body>\n<p>Formatted like: %s and %s.</p>\n</body>\n</html>", obj[0], obj[1])
	w.Write([]byte(output))
}

func main() {
	http.Handle("/xss-01/BenchmarkTest00711", &BenchmarkTest00711{})
	http.ListenAndServe(":8080", nil)
}
