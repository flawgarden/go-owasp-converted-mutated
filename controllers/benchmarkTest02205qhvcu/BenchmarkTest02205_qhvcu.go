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
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/queue.tmt with name queue_is_empty_1_negative 
//Used extensions: MACRO_Create_Queue -> ~[MACRO_QueueName]~ := list.New() | MACRO_Add_Fixed_VAR_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[VAR_~[TYPE@1]~@1]~) | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231
//Program:
package controllers

import (
"fmt"
"net/http"
"os"
"strings"
"container/list"
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

queue787231 := list.New()
queue787231.PushBack(bar)
value7843 := "ElVRh"
if queue787231.Len() > 0 {
    value7843 = queue787231.Front().Value.(string)
    queue787231.Remove(queue787231.Front())
}
bar = value7843

	fileName := ""
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
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
