package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type BenchmarkTest00039 struct{}

func (bt BenchmarkTest00039) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var param string
	flag := true

	for name, values := range r.Form {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest00039" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	sqlStr := fmt.Sprintf("SELECT userid FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", param)
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/goseccode")
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	results, err := db.Query(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	defer results.Close()

	var userID string
	if results.Next() {
		if err := results.Scan(&userID); err != nil {
			http.Error(w, "Error retrieving results", http.StatusInternalServerError)
			return
		}
	}

	response := map[string]string{"userid": userID}
	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
