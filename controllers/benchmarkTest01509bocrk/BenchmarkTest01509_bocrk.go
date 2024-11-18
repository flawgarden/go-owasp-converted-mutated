//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: []
//-------------
//Semgrep analysis results: [79]
//CodeQL analysis results: []
//Gosec analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01509/BenchmarkTest01509.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/concurrency/channels.tmt with name channel_struct_positive 
//Used extensions: 
//Program:
package controllers

import (
"net/http"
"strings"
"xorm.io/xorm"
"sync"
)

type BenchmarkTest01509 struct {
	engine *xorm.Engine
}

func (b *BenchmarkTest01509) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest01509) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01509")
	if param == "" {
		param = ""
	}

dataChannel := make(chan Data, 1)
dataChannel <- Data{Value: param}

var wg sync.WaitGroup
wg.Add(1)

go func() {
    data := <-dataChannel
    data.Value += "suffix"
    dataChannel <- data
}()

wg.Wait()

readData := <-dataChannel
param = readData.Value

	bar := new(Test).doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := ""
	if param != "" {
		bar = strings.Split(param, " ")[0]
	}
	return bar
}
