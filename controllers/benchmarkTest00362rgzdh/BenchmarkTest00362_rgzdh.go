//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: []
//-------------
//Gosec analysis results: [676, 22, 703]
//CodeQL analysis results: [563]
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00362/BenchmarkTest00362.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/assignment.tmt with name lazy_eval_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type BenchmarkTest00362 struct{}

func (b *BenchmarkTest00362) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest00362")

lazyValue := func() string { return "" }
param = lazyValue()

	if param == "" {
		param = ""
	}

	bar := param
	fileName := filepath.Join("testfiles", bar)
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fos, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}

	_, _ = w.Write([]byte("Now ready to write to file: " + fileName))
}

func main() {
	http.Handle("/pathtraver-00/BenchmarkTest00362", &BenchmarkTest00362{})
	http.ListenAndServe(":8080", nil)
}
