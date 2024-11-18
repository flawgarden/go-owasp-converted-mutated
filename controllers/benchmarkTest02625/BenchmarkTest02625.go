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

type BenchmarkTest02625 struct{}

func (b *BenchmarkTest02625) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02625="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramval)
		if paramLoc < 0 {
			http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02625"), http.StatusBadRequest)
			return
		}
	}
	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := paramLoc + len(paramval) + len(param)
	if ampersandLoc < len(queryString) {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("CALL %s", bar)

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

	user := models.User{}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func doSomething(param string) string {
	bar := "safe!"
	map82391 := make(map[string]interface{})
	map82391["keyA-82391"] = "a-Value"
	map82391["keyB-82391"] = param
	map82391["keyC"] = "another-Value"
	bar = map82391["keyB-82391"].(string)
	return bar
}
