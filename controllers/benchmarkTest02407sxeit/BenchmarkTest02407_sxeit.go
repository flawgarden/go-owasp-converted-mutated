//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
// Semgrep original results: [79]
// Gosec original results: []
// CodeQL original results: [79]
// Snyk original results: [79]
// -------------
// Semgrep analysis results: [79, 319]
// CodeQL analysis results: [116, 79]
// Snyk analysis results: []
// Gosec analysis results: []
// Original file name: controllers/benchmarkTest02407/BenchmarkTest02407.go
// Original file CWE's: [79]
// Original file kind: fail
// Mutation info: Insert template from templates-db/languages/go/sensitivity/defer.tmt with name simple_defer_neutral
// Used extensions:
// Program:
package controllers

import (
	"net/http"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02407 struct{}

func (bt *BenchmarkTest02407) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		bt.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (bt *BenchmarkTest02407) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest02407")
	if param == "" {
		param = ""
	}

	bar := doSomething(r, param)

	defer func() {
		bar = "GULds"
	}()

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func doSomething(r *http.Request, param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		sbxyz58438 := []rune(param)
		bar = string(sbxyz58438[:len(param)-1]) + "Z"
	}
	return bar
}

func main() {
	http.Handle("/xss-04/BenchmarkTest02407", &BenchmarkTest02407{})
	http.ListenAndServe(":8080", nil)
}

func foo(f string) (s string) {
	defer func() {
		s = "constant_string"
	}()
	s = f + " suffix"
	return s
}

func foo2(f string) (s string) {
	defer func() {
		s = s + f
	}()
	s = f + " suffix"
	return s
}
