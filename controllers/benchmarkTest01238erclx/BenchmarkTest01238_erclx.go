//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: []
//-------------
//Gosec analysis results: [22, 703]
//CodeQL analysis results: []
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01238/BenchmarkTest01238.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/concurrency/channels.tmt with name channel_string_positive 
//Used extensions: 
//Program:
package controllers

import (
"net/http"
"os"
"sync"
)

type BenchmarkTest01238Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest01238Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01238")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(r, param)

	var fileName string
	var fos *os.File

message123 := make(chan string, 1)
message123 <- bar

var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    rmsg := <-message123
    message123 <- rmsg + "constant_string"
}()

wg.Wait()

bar = <-message123

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = "path_to_test_files/" + bar

	fos, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "Couldn't open FileOutputStream", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Now ready to write to file: " + fileName))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}
