//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: []
//-------------
//Gosec analysis results: [22, 276, 703]
//CodeQL analysis results: []
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest02034/BenchmarkTest02034.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/virtuality/default.tmt with name binary_op_interface_default2_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest02034 struct{}

func (b *BenchmarkTest02034) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var param string
	if headers := r.Header["BenchmarkTest02034"]; len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)
	bar := doSomething(r, param)

	fileName := "testfiles/" + bar

var a12341 BinaryOpInterfaceDefault2 = &BinaryOpInterfaceDefault2Implementation{}
fileName = a12341.InterfaceCall(fileName, "")

	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte("Now ready to write to file: " + htmlEncode(fileName)))
}

func doSomething(r *http.Request, param string) string {
	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}

func htmlEncode(str string) string {
	// Dummy implementation for HTML encoding - replace with proper encoding if necessary
	return str
}
