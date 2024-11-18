package controllers

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type BenchmarkTest02107 struct{}

func (b *BenchmarkTest02107) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest02107")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	fileTarget := bar
	fmt.Fprintf(w, "Access to file: '%s' created.\n", htmlEscape(fileTarget))
	if _, err := os.Stat(fileTarget); err == nil {
		fmt.Fprintln(w, "And file already exists.")
	} else {
		fmt.Fprintln(w, "But file doesn't exist yet.")
	}
}

func doSomething(param string) string {
	bar := ""
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func htmlEscape(s string) string {
	return template.HTMLEscapeString(s)
}

func main() {
	http.Handle("/pathtraver-02/BenchmarkTest02107", &BenchmarkTest02107{})
	http.ListenAndServe(":8080", nil)
}
