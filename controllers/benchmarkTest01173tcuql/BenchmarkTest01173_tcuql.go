//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79]
//CodeQL analysis results: []
//Snyk analysis results: [79]
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest01173/BenchmarkTest01173.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/concurrency/channels.tmt with name channel_struct_data_int_positive 
//Used extensions: 
//Program:
package controllers

import (
"net/http"
"net/url"
_ "github.com/go-sql-driver/mysql"
"sync"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01173 struct{}

func (bt *BenchmarkTest01173) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Get("X-XSS-Protection")
	param := r.Header.Get("Referer")

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")

dataChannel := make(chan DataInt, 1)
dataChannel <- DataInt{Value: 1}

var wg sync.WaitGroup
wg.Add(1)

go func() {
    defer wg.Done()
    data := <-dataChannel
    data.Value += 1
    dataChannel <- data
}()

wg.Wait()

result := <-dataChannel
if (result.Value == 1) {
    bar = "constant_string"
}

	w.Write([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value

		bar = valuesList[0] // get the param value
	}
	return bar
}

func init() {
	http.Handle("/xss-02/BenchmarkTest01173", &BenchmarkTest01173{})
}
