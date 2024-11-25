//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: []
//-------------
//Gosec analysis results: [22, 276, 703]
//CodeQL analysis results: [563]
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest02034/BenchmarkTest02034.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/queue.tmt with name queue_poll_all_positive 
//Used extensions: MACRO_Create_Queue -> ~[MACRO_QueueName]~ := list.New() | MACRO_Add_EXPR_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[EXPR_~[TYPE@1]~]~) | MACRO_Add_EXPR_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[EXPR_~[TYPE@1]~]~) | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231
//Program:
package controllers

import (
"fmt"
"net/http"
"net/url"
"os"
"container/list"
)

type BenchmarkTest02034 struct{}

func (b *BenchmarkTest02034) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var param string
	if headers := r.Header["BenchmarkTest02034"]; len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)
	bar := doSomething(r, param)

queue787231 := list.New()
queue787231.PushBack("WDXHz")
queue787231.PushBack("xQKiS")
queue787231.PushBack(bar)
value7847 := "hJHkI"
for queue787231.Len() > 0 {
    element := queue787231.Front()
    if element != nil {
        value7847 = element.Value.(string)
        queue787231.Remove(element)
    }
}
bar = value7847

	fileName := "testfiles/" + bar

	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte("Now ready to write to file: " + htmlEncode(fileName)))
}

func doSomething(r *http.Request, param string) string {
	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}

func htmlEncode(str string) string {
	// Dummy implementation for HTML encoding - replace with proper encoding if necessary
	return str
}