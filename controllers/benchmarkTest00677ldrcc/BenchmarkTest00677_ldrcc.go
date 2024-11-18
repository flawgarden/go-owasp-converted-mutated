//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: [89]
//-------------
//Semgrep analysis results: [89, 79, 319]
//Gosec analysis results: [676, 89, 703]
//CodeQL analysis results: [563]
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00677/BenchmarkTest00677.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/regularField.tmt with name class_with_string_field_negative 
//Used extensions: 
//Program:
package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00677 struct{}

func (b *BenchmarkTest00677) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest00677")
	if param == "" {
		param = ""
	}
	bar := ""

	if param != "" {
		bar = string(param)
	}

sh := NewStringHolder()
sh.value = ""
bar = sh.value

	sqlQuery := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	results, err := db.Query(sqlQuery)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	defer results.Close()

	w.Write([]byte("Your results are: <br>"))
	var username string
	for results.Next() {
		err := results.Scan(&username)
		if err != nil {
			http.Error(w, "Error processing query", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(fmt.Sprintf("%s<br>", username)))
	}

	if err = results.Err(); err != nil {
		http.Error(w, "Error processing query", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.Handle("/sqli-01/BenchmarkTest00677", &BenchmarkTest00677{})
	http.ListenAndServe(":8080", nil)
}
