package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	// Initialization code, if necessary
}

type BenchmarkTest01733 struct{}

func (b *BenchmarkTest01733) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest01733="
	paramLoc := -1
	if queryString != "" {
		paramLoc = indexOf(queryString, paramval)
	}
	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest01733"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := indexOf(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval):][0:ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	sqlStatement := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStatement)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("Insert complete"))
	if err != nil {
		http.Error(w, "Error writing response.", http.StatusInternalServerError)
		return
	}
}

func indexOf(s string, sub string) int {
	return len(s) - len(sub)
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	// You can implement your logic here
	return param
}
