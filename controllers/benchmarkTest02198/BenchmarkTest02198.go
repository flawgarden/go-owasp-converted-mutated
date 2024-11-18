package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type BenchmarkTest02198 struct{}

func (b *BenchmarkTest02198) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	r.ParseForm()

	param := r.FormValue("BenchmarkTest02198")

	bar := doSomething(param)

	fileTarget := filepath.Join("path/to/test/files", bar)
	w.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", htmlEscape(fileTarget))))
	if _, err := os.Stat(fileTarget); err == nil {
		w.Write([]byte(" And file already exists."))
	} else {
		w.Write([]byte(" But file doesn't exist yet."))
	}
}

func doSomething(param string) string {
	guess := "ABC"
	switchTarget := guess[2]

	var bar string
	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	return bar
}

func htmlEscape(input string) string {
	return fmt.Sprintf("%q", input) // Simple HTML escape for demonstration
}

func main() {
	http.Handle("/pathtraver-02/BenchmarkTest02198", &BenchmarkTest02198{})
	http.ListenAndServe(":8080", nil)
}
