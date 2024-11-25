//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: [22]
//-------------
//Gosec analysis results: [676, 22, 276, 703]
//CodeQL analysis results: [22, 36, 73, 99, 23, 116, 79]
//Snyk analysis results: []
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest00363/BenchmarkTest00363.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/queue.tmt with name queue_contains_1_positive 
//Used extensions: MACRO_Create_Queue -> ~[MACRO_QueueName]~ := list.New() | MACRO_Add_EXPR_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[EXPR_~[TYPE@1]~]~) | MACRO_Add_Fixed_VAR_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[VAR_~[TYPE@1]~@1]~) | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231
//Program:
package controllers

import (
"encoding/json"
"fmt"
"net/http"
"os"
"container/list"
)

type BenchmarkTest00363 struct{}

func (b BenchmarkTest00363) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest00363")
	if param == "" {
		param = ""
	}
	bar := param

queue787231 := list.New()
queue787231.PushBack("ZUjXl")
queue787231.PushBack(bar)
value7845 := "XkcCe"
for e := queue787231.Front(); e != nil; e = e.Next() {
    if e.Value == bar {
        value7845 = e.Value.(string)
        break
    }
}
bar = value7845

	fileName := "testfiles/" + bar

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintln(w, "Couldn't open FileOutputStream on file: '"+fileName+"'")
		return
	}
	defer f.Close()

	response := map[string]string{
		"message": "Now ready to write to file: " + escapeHTML(fileName),
	}
	output, _ := json.Marshal(response)
	w.Write(output)
}

func escapeHTML(input string) string {
	return input // реализация для экранирования HTML
}

func main() {
	http.Handle("/pathtraver-00/BenchmarkTest00363", BenchmarkTest00363{})
	http.ListenAndServe(":8080", nil)
}
