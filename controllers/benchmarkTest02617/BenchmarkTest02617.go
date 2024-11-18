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

type BenchmarkTest struct{}

func (b *BenchmarkTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	paramVal := "BenchmarkTest02617"
	param := queryString.Get(paramVal)

	if param == "" {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02617' in query string.", http.StatusBadRequest)
		return
	}

	param = urlDecode(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", param)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Query execution error", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "JSON marshal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func urlDecode(param string) string {
	result, _ := url.QueryUnescape(param)
	return result
}

func doSomething(r *http.Request, param string) string {
	// Implementation of some functionality
	return param + " SafeStuff"
}
