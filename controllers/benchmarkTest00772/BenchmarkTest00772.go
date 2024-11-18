package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00772 struct{}

func (b *BenchmarkTest00772) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := r.URL.Query()["BenchmarkTest00772"]
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	var bar string
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

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

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
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
		row := make([]interface{}, len(columns))
		for i := range row {
			row[i] = new(sql.RawBytes)
		}
		if err := rows.Scan(row...); err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}
		result := make(map[string]interface{})
		for i, col := range columns {
			val := row[i].(*sql.RawBytes)
			result[col] = string(*val)
		}
		results = append(results, result)
	}
	output, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func main() {
	http.Handle("/sqli-01/BenchmarkTest00772", &BenchmarkTest00772{})
	http.ListenAndServe(":8080", nil)
}
