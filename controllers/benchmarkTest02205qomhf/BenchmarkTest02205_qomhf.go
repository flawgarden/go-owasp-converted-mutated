//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: []
//CodeQL original results: [22]
//Snyk original results: [22]
//-------------
//CodeQL analysis results: [22, 36, 73, 99, 23, 116, 79, 117]
//Snyk analysis results: []
//Semgrep analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest02205/BenchmarkTest02205.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/concurrency/concurrency.tmt with name thread_set_positive 
//Used extensions: 
//Program:
package controllers

import (
"fmt"
"net/http"
"os"
"strings"
"sync"
)

type BenchmarkTest02205 struct{}

func (bt *BenchmarkTest02205) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	param := r.FormValue("BenchmarkTest02205")

	bar := doSomething(param)

	fileName := ""
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}

w := &Wrapper[string]{Value: bar}
task := NewSettingTask(w, bar)
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    task.Run()
}()
wg.Wait()
bar = w.Value

	}()

	fileName = "testfiles/" + bar

	fos, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}

	response := fmt.Sprintf("Now ready to write to file: %s", htmlspecialchars(fileName))
	w.Write([]byte(response))
}

func doSomething(param string) string {
	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}

func htmlspecialchars(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}