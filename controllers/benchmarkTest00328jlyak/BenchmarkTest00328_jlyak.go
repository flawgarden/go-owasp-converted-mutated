//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [89]
//Gosec original results: []
//CodeQL original results: [89]
//Snyk original results: []
//-------------
//Semgrep analysis results: [89, 79]
//CodeQL analysis results: [116, 79]
//Gosec analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00328/BenchmarkTest00328.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/virtuality/class.tmt with name derived_binary_op1_negative 
//Used extensions: 
//Program:
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

var a12341 BaseBinaryOpClass = &DerivedBinaryOpClass1{}
sqlStr = a12341.VirtualCall("", sqlStr)

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
