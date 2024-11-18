package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00328 struct{}

func (b *BenchmarkTest00328) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if headers := r.Header["BenchmarkTest00328"]; len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := param
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	}

	sqlStr := fmt.Sprintf("{call %s}", bar)

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

	var results []map[string]interface{}
	columns, _ := rows.Columns()
	for rows.Next() {
		values := make([]interface{}, len(columns))
		for i := range values {
			values[i] = new(sql.RawBytes)
		}
		if err := rows.Scan(values...); err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}
		rowData := make(map[string]interface{})
		for i, col := range columns {
			rowData[col] = string(*values[i].(*sql.RawBytes))
		}
		results = append(results, rowData)
	}

	output, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}
