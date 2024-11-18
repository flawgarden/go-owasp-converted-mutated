//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: []
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79, 319]
//Snyk analysis results: []
//Gosec analysis results: []
//CodeQL analysis results: []
//Original file name: controllers/benchmarkTest02332/BenchmarkTest02332.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/list.tmt with name list_set_negative 
//Used extensions: MACRO_Create_List -> ~[MACRO_ListName]~ := make([] ~[TYPE@1]~, 0) | MACRO_Add_CONST_ToList -> ~[MACRO_ListName]~ = append(~[MACRO_ListName]~, ~[CONST_~[TYPE@1]~@1]~) | MACRO_ListName -> list787231 | MACRO_ListName -> list787231 | MACRO_ListName -> list787231 | MACRO_ListName -> list787231
//Program:
package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest02332 struct{}

func (b *BenchmarkTest02332) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	flag := true
	for name, values := range r.URL.Query() {

list787231 := make([] string, 0)
list787231 = append(list787231, "SbyMC")
name = list787231[0]

		if flag {
			for _, value := range values {
				if value == "BenchmarkTest02332" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func doSomething(r *http.Request, param string) string {
	bar := ""
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}
	return bar
}

func main() {
	http.Handle("/xss-04/BenchmarkTest02332", &BenchmarkTest02332{})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
