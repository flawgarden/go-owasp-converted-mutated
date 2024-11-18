package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00034 struct{}

func (b *BenchmarkTest00034) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("BenchmarkTest00034")
	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		http.Error(w, "Error executing query", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []map[string]interface{}
	columns, _ := rows.Columns()
	for rows.Next() {
		columnValues := make([]interface{}, len(columns))
		for i := range columnValues {
			columnValues[i] = new(interface{})
		}
		if err := rows.Scan(columnValues...); err != nil {
			http.Error(w, "Error scanning results", http.StatusInternalServerError)
			return
		}

		rowMap := make(map[string]interface{})
		for i, col := range columnValues {
			rowMap[columns[i]] = *(col.(*interface{}))
		}
		results = append(results, rowMap)
	}

	w.Header().Set("Content-Type", "application/json")
	output, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}
