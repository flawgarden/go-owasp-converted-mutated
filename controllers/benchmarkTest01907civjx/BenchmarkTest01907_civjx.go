//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: []
//-------------
//Gosec analysis results: [22, 703]
//CodeQL analysis results: [570, 571, 563]
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01907/BenchmarkTest01907.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/conditional/conditionswitch.tmt with name switch_fallthrough_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest01907 struct{}

func (b *BenchmarkTest01907) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01907")
	param, _ = url.QueryUnescape(param)

	bar := doSomething(r, param)

	var fileName string

value := 0
switch {
case value < 0:
    bar = "fixed_string"
case value == 0:
    bar = bar + "_suffix"
    fallthrough
case value > 0:
    bar = "fixed_string"
default:
    bar = "Unknown"
}

	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = "testfiles/" + bar

	fos, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}
	fmt.Fprintln(w, "Now ready to write to file: "+fileName)
}

func doSomething(r *http.Request, param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
}
