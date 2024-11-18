//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: [78]
//-------------
//Gosec analysis results: [78, 703]
//CodeQL analysis results: [116, 79, 563]
//Snyk analysis results: []
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest01360/BenchmarkTest01360.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/list.tmt with name list_clear_start_negative 
//Used extensions: MACRO_Create_List -> ~[MACRO_ListName]~ := make([] ~[TYPE@1]~, 0) | MACRO_Add_VAR_ToList -> ~[MACRO_ListName]~ = append(~[MACRO_ListName]~, ~[VAR_~[TYPE@1]~@1]~) | MACRO_ListName -> list787231 | MACRO_ListName -> list787231 | MACRO_ListName -> list787231 | MACRO_ListName -> list787231 | MACRO_ListName -> list787231
//Program:
package controllers

import (
	"net/http"
	"os"
	"os/exec"
)

type BenchmarkTest01360Controller struct{}

func (c *BenchmarkTest01360Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01360")

list787231 := make([] string, 0)
list787231 = append(list787231, param)
list787231 = nil
param = list787231[0]

	bar := new(Test).doSomething(r, param)

	cmd := ""
	if isWindows() {
		cmd = "cmd /c echo "
	}

	output, err := exec.Command(cmd + bar).Output()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func isWindows() bool {
	return len(os.Getenv("OS")) > 0 && os.Getenv("OS") == "Windows_NT"
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
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
