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
//Original file name: controllers/benchmarkTest02635/BenchmarkTest02635.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/constructors.tmt with name class_with_nested_string_initialization_positive 
//Used extensions: 
//Program:
package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02635 struct {
}

func (bt *BenchmarkTest02635) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02635="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02635"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

nsh := NewNestedStringHolder(bar)
bar = nsh.GetValue()

	sqlStr := "SELECT * from USERS where USERNAME=? and PASSWORD='" + bar + "'"

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

	var user models.User
	err = statement.QueryRow("foo").Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

func doSomething(param string) string {
	bar := ""
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

	return bar
}
