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

type BenchmarkTest02178 struct{}

func (b *BenchmarkTest02178) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest02178")
	if param == "" {
		param = ""
	}
	bar := doSomething(param)
	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			http.Error(w, "Error scanning results", http.StatusInternalServerError)
			return
		}
		results = append(results, user)
	}

	output, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error marshalling results", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func doSomething(param string) string {
	bar := "safe!"
	map38026 := make(map[string]string)
	map38026["keyA-38026"] = "a-Value"
	map38026["keyB-38026"] = param
	map38026["keyC"] = "another-Value"
	bar = map38026["keyB-38026"]
	return bar
}
