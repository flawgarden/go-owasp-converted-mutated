package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01713 struct{}

func (b *BenchmarkTest01713) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("BenchmarkTest01713")

	if param == "" {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest01713' in query string.", http.StatusBadRequest)
		return
	}

	bar := b.doSomething(param)
	sqlStr := fmt.Sprintf("{call %s}", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error opening database connection.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			http.Error(w, "Error scanning rows.", http.StatusInternalServerError)
			return
		}
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error marshaling response.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (b *BenchmarkTest01713) doSomething(param string) string {
	bar := "safe!"
	map83939 := make(map[string]interface{})
	map83939["keyA-83939"] = "a_Value"
	map83939["keyB-83939"] = param
	map83939["keyC"] = "another_Value"
	bar = map83939["keyB-83939"].(string)
	bar = map83939["keyA-83939"].(string)
	return bar
}
