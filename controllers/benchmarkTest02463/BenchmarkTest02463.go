package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest02463 struct{}

func (bt *BenchmarkTest02463) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := r.URL.Query()["BenchmarkTest02463"]
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	bar := doSomething(param)

	fileTarget := fmt.Sprintf("%s/%s", os.Getenv("TESTFILES_DIR"), bar)
	w.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", htmlEscape(fileTarget))))
	if _, err := os.Stat(fileTarget); err == nil {
		w.Write([]byte(" And file already exists."))
	} else {
		w.Write([]byte(" But file doesn't exist yet."))
	}
}

func doSomething(param string) string {
	bar := ""

	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	return bar
}

func htmlEscape(str string) string {
	return str // Replace with actual HTML escaping logic if necessary
}
