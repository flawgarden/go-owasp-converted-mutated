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

func init() {
	// Initialize database connection or any other setup here
}

type BenchmarkTest02644 struct{}

func (b *BenchmarkTest02644) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.RawQuery
	paramVal := "BenchmarkTest02644="
	paramLoc := -1

	if queryString != "" {
		paramLoc = findParamLocation(queryString, paramVal)
	}
	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02644"), http.StatusBadRequest)
		return
	}

	param := extractParamValue(queryString, paramLoc, paramVal)
	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Database connection error.", http.StatusInternalServerError)
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
			http.Error(w, "Error reading results.", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	output, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error marshaling JSON.", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func findParamLocation(queryString, paramVal string) int {
	return strings.Index(queryString, paramVal)
}

func extractParamValue(queryString string, paramLoc int, paramVal string) string {
	param := queryString[paramLoc+len(paramVal):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramVal) : paramLoc+ampersandLoc]
	}
	val, _ := url.QueryUnescape(param)
	return val
}

func doSomething(param string) string {
	return param
}
