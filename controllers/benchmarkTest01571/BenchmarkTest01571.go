package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest01571 struct{}

func (b *BenchmarkTest01571) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	values := r.Form["BenchmarkTest01571"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := b.doSomething(r, param)

	fileTarget := fmt.Sprintf("%s/Test.txt", bar)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintf(w, "Access to file: '%s' created.", fileTarget)
	if _, err := os.Stat(fileTarget); err == nil {
		fmt.Fprintln(w, " And file already exists.")
	} else {
		fmt.Fprintln(w, " But file doesn't exist yet.")
	}
}

func (b *BenchmarkTest01571) doSomething(r *http.Request, param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[2]

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
