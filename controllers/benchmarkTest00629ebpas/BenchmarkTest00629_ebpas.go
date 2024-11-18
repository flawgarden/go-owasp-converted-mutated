//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: [22]
//-------------
//Gosec analysis results: [22, 703]
//CodeQL analysis results: [22, 36, 73, 99, 23, 116, 79]
//Snyk analysis results: []
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest00629/BenchmarkTest00629.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/concurrency/concurrency.tmt with name null_and_restore_condvar_positive 
//Used extensions: 
//Program:
package controllers

import (
"net/http"
"os"
"path/filepath"
"sync"
)

type BenchmarkTest00629 struct{}

func (b *BenchmarkTest00629) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00629")
	if param == "" {
		param = ""
	}

	var bar string
	num := 196

a := NewNullAndRestore(param)
var wg sync.WaitGroup
wg.Add(2)
go func() {
    defer wg.Done()
	a.NullMethod()
}()
go func() {
    defer wg.Done()
	a.Restore()
}()
wg.Wait()

param = a.Get()

	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	fileName := filepath.Join(os.TempDir(), bar)
	file, err := os.Open(fileName)
	if err != nil {
		http.Error(w, "Couldn't open InputStream on file: "+fileName, http.StatusInternalServerError)
		return
	}
	defer file.Close()

	bContent := make([]byte, 1000)
	size, err := file.Read(bContent)
	if err != nil {
		http.Error(w, "Problem getting InputStream: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte("The beginning of file: '" + fileName + "' is:\n\n"))
	w.Write(bContent[:size])
}
