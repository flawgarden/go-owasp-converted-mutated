//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: [89]
//-------------
//Semgrep analysis results: [89, 79]
//Gosec analysis results: [89, 703]
//CodeQL analysis results: [116, 79, 563]
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00027/BenchmarkTest00027.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/nested.tmt with name nested_field_simple_negative 
//Used extensions: 
//Program:
package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00027 struct{}

func (b *BenchmarkTest00027) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest00027) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.FormValue("BenchmarkTest00027")
	if param == "" {
		param = ""
	}

	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", param)

nested7231 := NewNestedFields1("VNpZN")
sqlStr = nested7231.nested1.value

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Update complete", "sql": sqlStr}
	output, _ := json.Marshal(response)
	w.Write(output)
}
