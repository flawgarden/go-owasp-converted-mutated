//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79]
//CodeQL analysis results: []
//Snyk analysis results: [79]
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest00822/BenchmarkTest00822.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/virtuality/default.tmt with name binary_op_interface_default_negative 
//Used extensions: 
//Program:
package controllers

import (
	"net/http"
	"strings"
)

type BenchmarkTest00822 struct{}

func (b *BenchmarkTest00822) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest00822="
	paramLoc := strings.Index(queryString, paramval)

	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest00822' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}

var a12341 BinaryOpInterfaceDefault = &BinaryOpInterfaceDefaultImplementation{}
param = a12341.InterfaceCall(param, param)

	bar := "This should never happen"
	num := 196

	if (500/42)+num > 200 {
		bar = param
	}

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte("Parameter value: " + bar))
}
