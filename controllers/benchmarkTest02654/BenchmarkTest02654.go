package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02654 struct{}

func (b *BenchmarkTest02654) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramVal := "BenchmarkTest02654="
	paramLoc := -1
	if queryString != "" {
		paramLoc = findParamLoc(queryString, paramVal)
	}

	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02654' in query string.", http.StatusBadRequest)
		return
	}

	param := extractParam(queryString, paramLoc, paramVal)

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("SQL executed successfully"))
}

func findParamLoc(queryString, paramVal string) int {
	return strings.Index(queryString, paramVal)
}

func extractParam(queryString string, paramLoc int, paramVal string) string {
	param := queryString[paramLoc+len(paramVal):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = param[:ampersandLoc]
	}
	decodedParam, _ := url.QueryUnescape(param)
	return decodedParam
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}
