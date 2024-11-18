package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00344 struct{}

func (b *BenchmarkTest00344) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest00344")
	param, _ = url.QueryUnescape(param)

	var bar string
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

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
	for rows.Next() {
		columns, _ := rows.Columns()
		values := make([]interface{}, len(columns))
		for i := range values {
			values[i] = new(sql.NullString)
		}
		if err := rows.Scan(values...); err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}
		row := make(map[string]interface{})
		for i, col := range columns {
			val := values[i].(*sql.NullString)
			row[col] = val.String
		}
		results = append(results, row)
	}

	output, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func main() {
	http.Handle("/sqli-00/BenchmarkTest00344", &BenchmarkTest00344{})
	http.ListenAndServe(":8080", nil)
}
