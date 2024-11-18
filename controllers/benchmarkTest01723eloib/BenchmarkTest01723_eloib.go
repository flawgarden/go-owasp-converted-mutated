//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: [89]
//-------------
//Semgrep analysis results: [89, 79]
//Gosec analysis results: [89]
//CodeQL analysis results: [563]
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01723/BenchmarkTest01723.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/queue.tmt with name queue_poll_all_positive 
//Used extensions: MACRO_Create_Queue -> ~[MACRO_QueueName]~ := list.New() | MACRO_Add_EXPR_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[EXPR_~[TYPE@1]~]~) | MACRO_Add_EXPR_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[EXPR_~[TYPE@1]~]~) | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | EXPR_string -> strings.ToLower(~[EXPR_string]~) | MACRO_QueueName -> queue787231
//Program:
package controllers

import (
"database/sql"
"fmt"
"net/http"
"net/url"
"strings"
_ "github.com/go-sql-driver/mysql"
"container/list"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01723 struct{}

func (b *BenchmarkTest01723) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery

queue787231 := list.New()
queue787231.PushBack(strings.ToLower("MyMwj"))
queue787231.PushBack("gSQBY")
value7847 := "kijDu"
for queue787231.Len() > 0 {
    element := queue787231.Front()
    if element != nil {
        value7847 = element.Value.(string)
        queue787231.Remove(element)
    }
}
queryString = value7847

	paramval := "BenchmarkTest01723="
	paramLoc := strings.Index(queryString, paramval)

	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest01723' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	sqlStr := fmt.Sprintf("SELECT userid FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var results int
	err = db.QueryRow(sqlStr).Scan(&results)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Your results are: %d", results)
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := param
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	}
	return bar
}
