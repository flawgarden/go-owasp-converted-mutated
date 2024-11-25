//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: [22]
//-------------
//Gosec analysis results: [676, 22, 703]
//CodeQL analysis results: [116, 79]
//Snyk analysis results: []
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest02561/BenchmarkTest02561.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/concurrency/channels.tmt with name channel_string_negative 
//Used extensions: 
//Program:
package controllers

import (
"fmt"
"net/http"
"os"
"path/filepath"
"sync"
)

type BenchmarkTest02561 struct{}

func (bt *BenchmarkTest02561) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.Query()
	param := queryString.Get("BenchmarkTest02561")
	if param == "" {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02561'", http.StatusBadRequest)
		return
	}

message123 := make(chan string, 1)
message123 <- param

var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    _ = <-message123
    message123 <- "constant_string"
}()

wg.Wait()

param = <-message123

	bar := doSomething(param)

	fileName := filepath.Join("path/to/testfiles", bar)
	fis, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Couldn't open FileInputStream on file: '%s'\n", fileName)
		return
	}
	defer fis.Close()

	b := make([]byte, 1000)
	size, _ := fis.Read(b)
	fmt.Fprintf(w, "The beginning of file: '%s' is:\n\n", fileName)
	w.Write(b[:size])
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}

func main() {
	http.Handle("/pathtraver-03/BenchmarkTest02561", &BenchmarkTest02561{})
	http.ListenAndServe(":8080", nil)
}
