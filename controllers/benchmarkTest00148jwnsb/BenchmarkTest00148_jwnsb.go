//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: []
//-------------
//Semgrep analysis results: [79]
//CodeQL analysis results: [563]
//Gosec analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00148/BenchmarkTest00148.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/virtuality/class.tmt with name base_binary_op_positive 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest00148Controller struct {
	http.Handler
}

func (c *BenchmarkTest00148Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("Referer") != "" {
		param = r.Header.Get("Referer")

var a12341 BaseBinaryOpClass
if 1064922805 != 1386398031 {
    a12341 = &DerivedBinaryOpClass1{}
} else {
    a12341 = &DerivedBinaryOpClass2{}
}
param = a12341.VirtualCall("", "")

	}

	param, _ = url.QueryUnescape(param)

	var bar string

	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", bar}
	_, _ = fmt.Fprintf(w, "Formatted like: %1$s and %2$s.", obj)
}
