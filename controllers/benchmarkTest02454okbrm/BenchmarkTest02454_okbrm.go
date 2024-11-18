//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: [89]
//-------------
//Semgrep analysis results: [79]
//Gosec analysis results: [89, 703]
//CodeQL analysis results: [89, 116, 79]
//Snyk analysis results: [89]
//Original file name: controllers/benchmarkTest02454/BenchmarkTest02454.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/lambdas/mutation.tmt with name unary_lambda_mutation_positive 
//Used extensions: 
//Program:
package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02454 struct{}

func (b *BenchmarkTest02454) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest02454")

s23423 := param
a12341 := &StringHolder{}
lambda1231 := func(s *StringHolder) {
    s.value = s23423
}
lambda1231(a12341)
param = a12341.value

	if param == "" {
		param = ""
	}

	bar := doSomething(r, param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	defer statement.Close()
	_, err = statement.Exec()
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal("Query executed successfully")
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func doSomething(r *http.Request, param string) string {
	return param
}
