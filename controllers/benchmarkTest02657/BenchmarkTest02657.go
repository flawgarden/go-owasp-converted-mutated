package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02657 struct {
}

func (b *BenchmarkTest02657) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02657="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramval)
	}

	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02657' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]

	ampersandLoc := paramLoc + len(paramval) + len(param)
	if ampersandLoc < len(queryString) && queryString[ampersandLoc] == '&' {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error connecting to database.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func doSomething(param string) string {
	num := 106
	if (7*18)+num > 200 {
		return "This_should_always_happen"
	}
	return param
}
