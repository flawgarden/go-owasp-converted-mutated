//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: []
//-------------
//Gosec analysis results: [22, 703]
//CodeQL analysis results: []
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest02112/BenchmarkTest02112.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/types/embedding.tmt with name embedding_anonymous_fields_negative 
//Used extensions: 
//Program:
package controllers

import (
	"net/http"
	"os"
)

type BenchmarkTest02112 struct{}

func (b *BenchmarkTest02112) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest02112")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	var fileName string
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

a := AnonFieldHolder{
    Value2: "const_str",
    Anon: Anon{
        Value1: bar,
    },
}
bar = a.Value2

	fileName = "testfiles/" + bar
	fos, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "Couldn't open FileOutputStream on file: '"+fileName+"'", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Now ready to write to file: " + fileName))
}

func doSomething(param string) string {
	guess := "ABC"
	switchTarget := guess[2]

	var bar string
	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	return bar
}
