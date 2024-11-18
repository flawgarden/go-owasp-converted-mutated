//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: [78]
//-------------
//Gosec analysis results: [78, 703]
//CodeQL analysis results: [116, 79, 570, 571, 563]
//Snyk analysis results: [79, 78]
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest01360/BenchmarkTest01360.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/conditional/conditionswitch.tmt with name switch_fallthrough_negative 
//Used extensions: 
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

value := 0
switch {
case value < 0:
    param = "fixed_string"
case value == 0:
    param = param + "_suffix"
    fallthrough
case value > 0:
    param = "fixed_string"
default:
    param = "Unknown"
}

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
