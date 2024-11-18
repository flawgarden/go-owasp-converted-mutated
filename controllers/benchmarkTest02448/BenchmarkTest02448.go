package controllers

import (
	"fmt"
	"net/http"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02448 struct{}

func (b *BenchmarkTest02448) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02448")
	if param == "" {
		param = ""
	}
	bar := doSomething(r, param)

	// Simulating session storage
	r.Header.Set(bar, "10340")

	w.Write([]byte(fmt.Sprintf("Item: '%s' with value: '10340' saved in session.", bar)))
}

func doSomething(r *http.Request, param string) string {
	bar := ""
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

func main() {
	http.Handle("/trustbound-01/BenchmarkTest02448", &BenchmarkTest02448{})
	http.ListenAndServe(":8080", nil)
}
