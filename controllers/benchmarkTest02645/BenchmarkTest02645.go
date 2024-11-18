package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02645 struct{}

func (b *BenchmarkTest02645) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query().Encode()
	paramval := "BenchmarkTest02645="
	paramLoc := strings.Index(queryString, paramval)

	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02645' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param = strings.TrimSpace(param)

	bar := doSomething(param)

	sqlStr := "SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='" + bar + "'"
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	results, err := db.Query(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer results.Close()

	var users []models.User
	for results.Next() {
		var user models.User
		if err := results.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			http.Error(w, "Error scanning results.", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	output, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func doSomething(param string) string {
	num := 196
	if (500/42)+num > 200 {
		return param
	}
	return "This should never happen"
}
