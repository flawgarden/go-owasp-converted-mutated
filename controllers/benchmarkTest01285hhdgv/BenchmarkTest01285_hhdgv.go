// PASS
// Semgrep original results: []
// Gosec original results: [78]
// CodeQL original results: [78]
// Snyk original results: [78]
// -------------
// Gosec analysis results: [676, 78, 703]
// CodeQL analysis results: [78, 116, 79]
// Snyk analysis results: []
// Semgrep analysis results: []
// Original file name: controllers/benchmarkTest01285/BenchmarkTest01285.go
// Original file CWE's: [78]
// Original file kind: fail
// Mutation info: Insert template from templates-db/languages/go/sensitivity/field/constructors.tmt with name class_with_array_initialization_neutral
// Used extensions: MACRO_Empty_string_Array -> []string{"", "", "", ""} | MACRO_Zero_Or_One -> 1
// Program:
package controllers

import (
	"net/http"
	"os/exec"
)

type BenchmarkTest01285 struct {
}

func (b *BenchmarkTest01285) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01285")
	if param == "" {
		param = ""
	}

	tmpArrayUnique42 := []string{"", "", "", ""}
	tmpArrayUnique42[0] = param
	ah := NewArrayHolderWithValues(tmpArrayUnique42)
	param = ah.Values[1]

	bar := b.doSomething(param)

	var cmd string
	osName := "UNIX"
	if osName == "Windows" {
		cmd = `echo `
	}

	command := exec.Command(cmd + bar)
	output, err := command.Output()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write(output)
}

func (b *BenchmarkTest01285) doSomething(param string) string {
	var bar string
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}

func main() {
	http.Handle("/cmdi-01/BenchmarkTest01285", &BenchmarkTest01285{})
	http.ListenAndServe(":8080", nil)
}
