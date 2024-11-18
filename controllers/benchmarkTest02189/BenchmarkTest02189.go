package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02189Controller struct{}

func (c *BenchmarkTest02189Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *BenchmarkTest02189Controller) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest02189")
	if param == "" {
		param = ""
	}
	bar := doSomething(r, param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM employees WHERE emplid='%s'", bar)
	rows, err := db.Query(sqlStr)
	if err != nil {
		http.Error(w, "Query execution error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var responseData []string
	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.EmplID, &emp.Name); err != nil {
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			return
		}
		responseData = append(responseData, emp.Name)
	}
	output, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func doSomething(r *http.Request, param string) string {
	return param
}

type Employee struct {
	EmplID string
	Name   string
}
