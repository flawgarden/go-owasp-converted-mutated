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

type BenchmarkTest02629 struct{}

func (b *BenchmarkTest02629) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02629="
	paramLoc := -1
	if queryString != "" {
		paramLoc = findParamLocation(queryString, paramval)
	}

	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02629' in query string.", http.StatusBadRequest)
		return
	}

	param := extractParamValue(queryString, paramLoc, paramval)

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

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func findParamLocation(queryString, paramval string) int {
	return strings.Index(queryString, paramval)
}

func extractParamValue(queryString string, paramLoc int, paramval string) string {
	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = param[:ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)
	return param
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}
	return bar
}
