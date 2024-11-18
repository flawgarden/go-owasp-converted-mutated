package controllers

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02632 struct{}

func (c *BenchmarkTest02632) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02632="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString)
	}
	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02632' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := len(queryString)
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}

	bar := doSomething(r, param)

	sqlQuery := "SELECT * FROM USERS WHERE USERNAME=? AND PASSWORD='" + bar + "'"

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error connecting to database.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare(sqlQuery)
	if err != nil {
		http.Error(w, "Error preparing SQL statement.", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec("foo")
	if err != nil {
		http.Error(w, "Error executing SQL statement.", http.StatusInternalServerError)
		return
	}

	// Process results (omitted for brevity)
}

func doSomething(r *http.Request, param string) string {
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

	return bar
}
