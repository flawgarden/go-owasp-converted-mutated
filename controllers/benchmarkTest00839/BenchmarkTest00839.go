package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00839 struct{}

func (b *BenchmarkTest00839) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest00839="
	paramLoc := -1

	if queryString != "" {
		paramLoc = findParamLoc(queryString, paramval)
	}

	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest00839' in query string.", http.StatusBadRequest)
		return
	}

	param := extractParam(queryString, paramLoc, paramval)

	param, _ = url.QueryUnescape(param)
	bar := calculateBar(param)
	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

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

	rows, err := statement.Query("foo")
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Username, &user.Password); err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	output, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func findParamLoc(queryString, paramval string) int {
	return strings.Index(queryString, paramval)
}

func extractParam(queryString string, paramLoc int, paramval string) string {
	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	return param
}

func calculateBar(param string) string {
	num := 106
	if (7*42 - num) > 200 {
		return "This should never happen"
	}
	return param
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
