//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: [78]
//-------------
//Gosec analysis results: [78, 703]
//CodeQL analysis results: [116, 79]
//Snyk analysis results: [79, 78]
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest01360/BenchmarkTest01360.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/virtuality/default.tmt with name binary_op_interface_default_negative 
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

var a12341 BinaryOpInterfaceDefault = &BinaryOpInterfaceDefaultImplementation{}
param = a12341.InterfaceCall(param, param)

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
