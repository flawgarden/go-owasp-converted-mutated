//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: [89]
//-------------
//Gosec analysis results: [89, 703]
//CodeQL analysis results: [89, 116, 79]
//Snyk analysis results: []
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest00192/BenchmarkTest00192.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/queue.tmt with name queue_is_empty_1_positive 
//Used extensions: MACRO_Create_Queue -> ~[MACRO_QueueName]~ := list.New() | MACRO_Add_Fixed_VAR_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[VAR_~[TYPE@1]~@1]~) | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231
//Program:
package controllers

import (
"database/sql"
"encoding/json"
"fmt"
"net/http"
"net/url"
_ "github.com/go-sql-driver/mysql"
"container/list"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00192 struct {
}

func (bt *BenchmarkTest00192) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest00192")

queue787231 := list.New()
queue787231.PushBack(param)
value7843 := "rucon"
if queue787231.Len() > 0 {
    value7843 = queue787231.Front().Value.(string)
}
param = value7843

	param, _ = url.QueryUnescape(param)

	var bar string
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query("foo")
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var username, password string
		rows.Scan(&username, &password)
		results = append(results, map[string]interface{}{"Username": username, "Password": password})
	}
	output, _ := json.Marshal(results)
	w.Write(output)
}
