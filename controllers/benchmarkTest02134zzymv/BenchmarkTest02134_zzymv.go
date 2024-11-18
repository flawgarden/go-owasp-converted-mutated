//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79, 319]
//CodeQL analysis results: [563]
//Snyk analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest02134/BenchmarkTest02134.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/list.tmt with name list_clear_start_negative 
//Used extensions: MACRO_Create_List -> ~[MACRO_ListName]~ := make([] ~[TYPE@1]~, 0) | MACRO_Add_VAR_ToList -> ~[MACRO_ListName]~ = append(~[MACRO_ListName]~, ~[VAR_~[TYPE@1]~@1]~) | MACRO_ListName -> list787231 | MACRO_ListName -> list787231 | MACRO_ListName -> list787231 | MACRO_ListName -> list787231 | MACRO_ListName -> list787231
//Program:
package controllers

import (
	"net/http"
)

type BenchmarkTest02134 struct{}

func (b *BenchmarkTest02134) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0")
	param := r.URL.Query().Get("BenchmarkTest02134")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

list787231 := make([] string, 0)
list787231 = append(list787231, bar)
list787231 = nil
bar = list787231[0]

	w.Write([]byte(bar))
}

func doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}

func main() {
	http.Handle("/xss-04/BenchmarkTest02134", &BenchmarkTest02134{})
	http.ListenAndServe(":8080", nil)
}
