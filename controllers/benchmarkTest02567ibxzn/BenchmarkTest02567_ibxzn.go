//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: [22]
//-------------
//Gosec analysis results: [22, 276]
//CodeQL analysis results: [22, 36, 73, 99, 23]
//Snyk analysis results: []
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest02567/BenchmarkTest02567.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/queue.tmt with name queue_remove_if_positive 
//Used extensions: MACRO_Create_Queue -> ~[MACRO_QueueName]~ := list.New() | MACRO_Add_Fixed_CONST_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[CONST_~[TYPE@1]~@1]~) | MACRO_Add_Fixed_VAR_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[VAR_~[TYPE@1]~@1]~) | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231
//Program:
package controllers

import (
"fmt"
"net/http"
"net/url"
"os"
"container/list"
)

const testFilesDir = "testfiles/"

type BenchmarkTest02567 struct{}

func (b *BenchmarkTest02567) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery

queue787231 := list.New()
queue787231.PushBack("YyVxj")
queue787231.PushBack(queryString)
for e := queue787231.Front(); e != nil; {
    next := e.Next()
    if e.Value == "YyVxj" {
        queue787231.Remove(e)
    }
    e = next
}
queryString = queue787231.Front().Value.(string)

	paramval := "BenchmarkTest02567="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramval) - 1
	}
	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02567' in query string.", http.StatusBadRequest)
		return
	}

	param, err := url.QueryUnescape(queryString[paramLoc+len(paramval):])
	if err != nil {
		http.Error(w, "Error decoding parameter", http.StatusBadRequest)
		return
	}

	bar := doSomething(param)

	fileName := testFilesDir + bar
	fos, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintf(w, "Couldn't open FileOutputStream on file: '%s'", fileName)
		return
	}
	defer fos.Close()

	fmt.Fprintf(w, "Now ready to write to file: %s", fileName)
}

func doSomething(param string) string {
	return param
}
