package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02550 struct{}

func (b *BenchmarkTest02550) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02550="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramval)
		if queryString[paramLoc:] != paramval {
			http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02550' in query string.", http.StatusBadRequest)
			return
		}
	}

	param := queryString[paramLoc+len(paramval):]

	bar := doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where username='%s'", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Error querying database", http.StatusInternalServerError)
		return
	}
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func doSomething(param string) string {
	// Пример экранирования или обработки параметра
	return param
}
