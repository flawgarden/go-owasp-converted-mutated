//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: [22]
//-------------
//Gosec analysis results: [676, 22, 276, 703]
//CodeQL analysis results: [116, 79]
//Snyk analysis results: [79, 23]
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest00363/BenchmarkTest00363.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/virtuality/interface.tmt with name impl_binary_op_interface_class2_negative 
//Used extensions: 
//Program:
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest00363 struct{}

func (b BenchmarkTest00363) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest00363")

var a12341 BinaryOpInterface = &ImplBinaryOpInterfaceClass2{}
param = a12341.InterfaceCall(param, "")

	if param == "" {
		param = ""
	}
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
