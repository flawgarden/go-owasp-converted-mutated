//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79]
//CodeQL analysis results: [563]
//Snyk analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest01173/BenchmarkTest01173.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/queue.tmt with name queue_poll_all_positive 
//Used extensions: MACRO_Create_Queue -> ~[MACRO_QueueName]~ := list.New() | MACRO_Add_EXPR_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[EXPR_~[TYPE@1]~]~) | MACRO_Add_EXPR_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[EXPR_~[TYPE@1]~]~) | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | EXPR_string -> ""
//Program:
package controllers

import (
"net/http"
"net/url"
_ "github.com/go-sql-driver/mysql"
"container/list"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01173 struct{}

func (bt *BenchmarkTest01173) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Get("X-XSS-Protection")
	param := r.Header.Get("Referer")

queue787231 := list.New()
queue787231.PushBack("SriRh")
queue787231.PushBack("")
value7847 := "NvOSv"
for queue787231.Len() > 0 {
    element := queue787231.Front()
    if element != nil {
        value7847 = element.Value.(string)
        queue787231.Remove(element)
    }
}
param = value7847

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
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
