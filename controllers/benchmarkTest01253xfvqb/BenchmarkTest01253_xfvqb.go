//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79]
//CodeQL analysis results: [116, 79]
//Snyk analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest01253/BenchmarkTest01253.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/queue.tmt with name queue_remove_if_positive 
//Used extensions: MACRO_Create_Queue -> ~[MACRO_QueueName]~ := list.New() | MACRO_Add_Fixed_CONST_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[CONST_~[TYPE@1]~@1]~) | MACRO_Add_Fixed_VAR_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[VAR_~[TYPE@1]~@1]~) | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231
//Program:
package controllers

import (
"fmt"
"net/http"
"container/list"
)

type BenchmarkTest01253 struct{}

func (b *BenchmarkTest01253) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0")
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01253")
	if param == "" {
		param = ""
	}

	bar := b.doSomething(param)

queue787231 := list.New()
queue787231.PushBack("WVlrE")
queue787231.PushBack(bar)
for e := queue787231.Front(); e != nil; {
    next := e.Next()
    if e.Value == "WVlrE" {
        queue787231.Remove(e)
    }
    e = next
}
bar = queue787231.Front().Value.(string)

	obj := []interface{}{"a", "b"}
	_, err := w.Write([]byte(fmt.Sprintf(bar, obj...)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (b *BenchmarkTest01253) doSomething(param string) string {
	bar := "safe!"
	map15481 := make(map[string]interface{})
	map15481["keyA-15481"] = "a-Value"
	map15481["keyB-15481"] = param
	map15481["keyC"] = "another-Value"
	bar = map15481["keyB-15481"].(string)

	return bar
}