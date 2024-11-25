//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79]
//CodeQL analysis results: []
//Snyk analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest00146/BenchmarkTest00146.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/anonstructures.tmt with name simple_anonymous_structure_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest00146 struct{}

func (b *BenchmarkTest00146) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("Referer")
	if param != "" {
		// URL Decode the header value
		param, _ = url.QueryUnescape(param)
	}

simpleStructure := struct {
    Field1 string
    Field2 string
}{
    Field1: param,
    Field2:  "sCOff",
}

param = simpleStructure.Field2

	bar := param
	if len(param) > 1 {
		sbxyz67327 := []rune(param)
		bar = string(append(sbxyz67327[:len(param)-1], 'Z'))
	}

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", "b"}
	_, _ = w.Write([]byte(fmt.Sprintf(bar, obj...)))
}

func createPoint(x, y string) struct {
    X string
    Y string
} {
    return struct {
        X string
        Y string
    }{
        X: x,
        Y: y,
    }
}

