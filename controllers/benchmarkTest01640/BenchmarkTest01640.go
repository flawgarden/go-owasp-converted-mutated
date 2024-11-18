package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func main() {
	http.HandleFunc("/crypto-02/BenchmarkTest01640", benchmarkTest01640)
	http.ListenAndServe(":8080", nil)
}

func benchmarkTest01640(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest01640="
	paramLoc := -1
	if queryString != "" {
		paramLoc = indexOf(queryString, paramval)
	}
	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest01640' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := indexOf(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}

	bar := doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where id=%s", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "JSON error", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func indexOf(s, substr string) int {
	for i := 0; i < len(s)-len(substr)+1; i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func doSomething(param string) string {
	guess := "ABC"
	switchTarget := guess[1]
	var bar string
	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}
	return bar
}
