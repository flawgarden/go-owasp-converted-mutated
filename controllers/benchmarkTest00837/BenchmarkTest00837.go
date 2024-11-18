package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00837 struct{}

func (b *BenchmarkTest00837) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	response := w
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.FormValue("BenchmarkTest00837")

	if param == "" {
		http.Error(response, "getQueryString() couldn't find expected parameter 'BenchmarkTest00837' in query string.", http.StatusBadRequest)
		return
	}

	bar := "safe!"
	bar = param

	sqlStr := fmt.Sprintf("{call %s}", bar)

	conn, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(response, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	rows, err := conn.Query(sqlStr)
	if err != nil {
		http.Error(response, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []map[string]interface{}
	cols, _ := rows.Columns()
	for rows.Next() {
		rowMap := make(map[string]interface{})
		columns := make([]interface{}, len(cols))
		for i := range columns {
			columns[i] = new(sql.RawBytes)
		}
		if err := rows.Scan(columns...); err != nil {
			http.Error(response, "Error processing request.", http.StatusInternalServerError)
			return
		}
		for i, col := range columns {
			rowMap[cols[i]] = string(*col.(*sql.RawBytes))
		}
		results = append(results, rowMap)
	}

	output, err := json.Marshal(results)
	if err != nil {
		http.Error(response, "Error processing request.", http.StatusInternalServerError)
		return
	}
	response.Write(output)
}

func main() {
	http.Handle("/sqli-01/BenchmarkTest00837", &BenchmarkTest00837{})
	http.ListenAndServe(":8080", nil)
}
