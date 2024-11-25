//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: [22]
//-------------
//Gosec analysis results: [676, 22, 276, 703]
//CodeQL analysis results: [22, 36, 73, 99, 23, 116, 79]
//Snyk analysis results: []
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest00363/BenchmarkTest00363.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/varargs/varargs.tmt with name varargs_get_string_with_index_positive 
//Used extensions: 
//Program:
package controllers

import (
"encoding/json"
"fmt"
"net/http"
"os"
"strings"
)

type BenchmarkTest00363 struct{}

func (b BenchmarkTest00363) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest00363")
	if param == "" {
		param = ""
	}

param = getStringWithIndex(0, param, "KkBpU")

	bar := param
	fileName := "testfiles/" + bar

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintln(w, "Couldn't open FileOutputStream on file: '"+fileName+"'")
		return
	}
	defer f.Close()

	response := map[string]string{
		"message": "Now ready to write to file: " + escapeHTML(fileName),
	}
	output, _ := json.Marshal(response)
	w.Write(output)
}

func escapeHTML(input string) string {
	return input // реализация для экранирования HTML
}

func main() {
	http.Handle("/pathtraver-00/BenchmarkTest00363", BenchmarkTest00363{})
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


