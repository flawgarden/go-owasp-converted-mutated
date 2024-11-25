//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: [22]
//-------------
//Gosec analysis results: [22, 276]
//CodeQL analysis results: [563]
//Snyk analysis results: []
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest02567/BenchmarkTest02567.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/constructors.tmt with name class_with_string_initialization_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

const testFilesDir = "testfiles/"

type BenchmarkTest02567 struct{}

func (b *BenchmarkTest02567) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02567="

sh := NewStringHolder()
queryString = sh.value

	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramval) - 1
	}
	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02567' in query string.", http.StatusBadRequest)
		return
	}

	param, err := url.QueryUnescape(queryString[paramLoc+len(paramval):])
	if err != nil {
		http.Error(w, "Error decoding parameter", http.StatusBadRequest)
		return
	}

	bar := doSomething(param)

	fileName := testFilesDir + bar
	fos, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintf(w, "Couldn't open FileOutputStream on file: '%s'", fileName)
		return
	}
	defer fos.Close()

	fmt.Fprintf(w, "Now ready to write to file: %s", fileName)
}

func doSomething(param string) string {
	return param
}
