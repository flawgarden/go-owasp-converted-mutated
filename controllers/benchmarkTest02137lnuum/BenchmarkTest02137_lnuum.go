//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: [78]
//-------------
//Gosec analysis results: [78, 703]
//CodeQL analysis results: [78]
//Snyk analysis results: []
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest02137/BenchmarkTest02137.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/constructors.tmt with name class_with_array_initialization_negative 
//Used extensions: MACRO_Empty_string_Array -> []string{"", "", "", ""}
//Program:
package controllers

import (
	"net/http"
	"os"
	"os/exec"
)

type BenchmarkTest02137 struct{}

func (b *BenchmarkTest02137) Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02137")

tmpArrayUnique42 := []string{"", "", "", ""}
tmpArrayUnique42[0] = param
ah := NewArrayHolderWithValues(tmpArrayUnique42)
param = ah.Values[1]

	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	var cmd []string
	if isWindows() {
		cmd = []string{"cmd.exe", "/c", "echo " + bar}
	} else {
		cmd = []string{"sh", "-c", "echo " + bar}
	}

	output, err := exec.Command(cmd[0], cmd[1:]...).Output()
	if err != nil {
		http.Error(w, "Error executing command", http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

func doSomething(param string) string {
	bar := ""
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}

func isWindows() bool {
	return os.Getenv("OS") == "Windows_NT"
}
