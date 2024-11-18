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

func init() {
	http.HandleFunc("/sqli-00/BenchmarkTest00342", BenchmarkTest00342)
}

func BenchmarkTest00342(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest00342")
	param, _ = url.QueryUnescape(param)

	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []map[string]interface{}
	columns, _ := rows.Columns()
	for rows.Next() {
		rowValues := make([]interface{}, len(columns))
		for i := range rowValues {
			rowValues[i] = new(sql.RawBytes)
		}

		if err := rows.Scan(rowValues...); err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}

		rowMap := make(map[string]interface{})
		for i, colName := range columns {
			val := rowValues[i].(*sql.RawBytes)
			rowMap[colName] = string(*val)
		}
		results = append(results, rowMap)
	}

	output, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

func main() {
	http.ListenAndServe(":8080", nil)
}
