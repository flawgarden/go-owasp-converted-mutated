//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: [89]
//-------------
//Semgrep analysis results: [89]
//Gosec analysis results: [89, 703]
//CodeQL analysis results: [89]
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01093/BenchmarkTest01093.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/queue.tmt with name queue_remove_if_positive 
//Used extensions: MACRO_Create_Queue -> ~[MACRO_QueueName]~ := list.New() | MACRO_Add_Fixed_CONST_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[CONST_~[TYPE@1]~@1]~) | MACRO_Add_Fixed_VAR_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[VAR_~[TYPE@1]~@1]~) | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231
//Program:
package controllers

import (
"database/sql"
"fmt"
"net/http"
"net/url"
_ "github.com/go-sql-driver/mysql"
"container/list"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01093 struct{}

func (b *BenchmarkTest01093) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01093")
	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

queue787231 := list.New()
queue787231.PushBack("QNIVj")
queue787231.PushBack(sqlStr)
for e := queue787231.Front(); e != nil; {
    next := e.Next()
    if e.Value == "QNIVj" {
        queue787231.Remove(e)
    }
    e = next
}
sqlStr = queue787231.Front().Value.(string)

	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	// Dummy response to mimic the original behavior
	w.Write([]byte("Query executed"))
}

func (b *BenchmarkTest01093) doSomething(param string) string {
	map18142 := map[string]interface{}{
		"keyA-18142": "a-Value",
		"keyB-18142": param,
		"keyC":       "another-Value",
	}
	return map18142["keyB-18142"].(string)
}
