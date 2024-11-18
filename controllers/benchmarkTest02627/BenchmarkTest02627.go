package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02627 struct {
}

func (bt *BenchmarkTest02627) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02627="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString)
	}
	if paramLoc == -1 {
		w.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest02627' in query string."))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := findAmpersand(queryString, paramLoc)
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)
	sqlStr := fmt.Sprintf("{call %s}", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}
	}
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func findAmpersand(queryString string, paramLoc int) int {
	return -1 // Implement find logic as needed
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		bar = string(param)
	}
	return bar
}
