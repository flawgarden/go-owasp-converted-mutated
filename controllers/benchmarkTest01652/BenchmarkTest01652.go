package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01652 struct{}

func (b *BenchmarkTest01652) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramVal := "BenchmarkTest01652="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramVal)
	}
	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest01652' in query string.", http.StatusBadRequest)
		return
	}
	param := queryString[paramLoc+len(paramVal):]
	ampersandLoc := paramLoc + len(paramVal)
	if ampersandLoc < len(queryString) {
		param = queryString[paramLoc+len(paramVal) : ampersandLoc]
	}
	param, err := url.QueryUnescape(param)
	if err != nil {
		http.Error(w, "Failed to decode parameter.", http.StatusBadRequest)
		return
	}
	bar := b.doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Database connection error.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Query execution error.", http.StatusInternalServerError)
		return
	}
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "JSON marshal error.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (b *BenchmarkTest01652) doSomething(param string) string {
	bar := "safe!"
	mapVar := make(map[string]interface{})
	mapVar["keyA-16429"] = "a_Value"
	mapVar["keyB-16429"] = param
	mapVar["keyC"] = "another_Value"
	bar = mapVar["keyB-16429"].(string)
	bar = mapVar["keyA-16429"].(string)
	return bar
}

func main() {
	http.Handle("/hash-01/BenchmarkTest01652", &BenchmarkTest01652{})
	http.ListenAndServe(":8080", nil)
}
