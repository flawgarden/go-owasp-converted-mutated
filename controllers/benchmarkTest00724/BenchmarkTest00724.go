package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest00724 struct{}

func (b *BenchmarkTest00724) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	values := r.Form["BenchmarkTest00724"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

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

	w.Header().Set("X-XSS-Protection", "0")
	fmt.Fprintln(w, bar)
}
