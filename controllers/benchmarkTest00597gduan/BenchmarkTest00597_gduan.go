//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: [89]
//-------------
//Semgrep analysis results: [89, 79]
//Gosec analysis results: [89, 703]
//CodeQL analysis results: [89, 116, 79]
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00597/BenchmarkTest00597.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/queue.tmt with name queue_is_empty_1_negative 
//Used extensions: MACRO_Create_Queue -> ~[MACRO_QueueName]~ := list.New() | MACRO_Add_Fixed_VAR_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[VAR_~[TYPE@1]~@1]~) | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231
//Program:
package controllers

import (
"database/sql"
"encoding/json"
"fmt"
"net/http"
"go-sec-code/models"
_ "github.com/go-sql-driver/mysql"
"container/list"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00597 struct{}

func (b *BenchmarkTest00597) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest00597) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	id := r.URL.Query().Get("id")

queue787231 := list.New()
queue787231.PushBack(id)
value7843 := "sAczJ"
if queue787231.Len() > 0 {
    value7843 = queue787231.Front().Value.(string)
    queue787231.Remove(queue787231.Front())
}
id = value7843

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", id)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(output)
}
