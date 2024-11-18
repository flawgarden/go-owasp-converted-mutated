//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79, 319]
//CodeQL analysis results: [563]
//Snyk analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest00477/BenchmarkTest00477.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/set.tmt with name set_add_simple_negative 
//Used extensions: MACRO_Create_Set -> ~[MACRO_SetName]~ := make(map[~[TYPE@1]~]struct{}) | MACRO_Add_Fixed_CONST_ToSet -> ~[MACRO_SetName]~[~[CONST_~[TYPE@1]~@1]~] = struct{}{} | MACRO_SetName -> set787231 | MACRO_SetName -> set787231 | MACRO_SetName -> set787231
//Program:
package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest00477 struct{}

func (b *BenchmarkTest00477) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00477")

set787231 := make(map[string]struct{})
set787231["BlZHk"] = struct{}{}
param = func() string {
    for k := range set787231 {
        return k
    }
    return "hhLQl"
}()

	bar := fmt.Sprintf("%s_SafeStuff", param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00477", &BenchmarkTest00477{})
	http.ListenAndServe(":8080", nil)
}
