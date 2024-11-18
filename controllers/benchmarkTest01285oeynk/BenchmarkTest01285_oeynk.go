//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: [78]
//-------------
//Gosec analysis results: [676, 78, 703]
//CodeQL analysis results: [116, 79, 563]
//Snyk analysis results: []
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest01285/BenchmarkTest01285.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/varargs/varargs.tmt with name varargs_get_first_string_from_array_negative 
//Used extensions: 
//Program:
package controllers

import (
"net/http"
"os/exec"
"strings"
)

type BenchmarkTest01285 struct {
}

func (b *BenchmarkTest01285) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01285")
	if param == "" {
		param = ""
	}

	bar := b.doSomething(param)

	var cmd string

bar = getFirstStringFromArray("ifcQK", "slfaS")

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

func getFirstString(lines ...string) string {
    return getStringWithIndex(0, lines...)
}

func getStringWithIndex(ind int, lines ...string) string {
    return lines[ind]
}

func getFirstStringFromArray(lines ...string) string {
    return lines[0]
}

func varargsWithGenerics[T any](elements ...T) T {
    return elements[0]
}

func combineStrings(strs ...string) string {
    return strings.Join(strs, ", ")
}


