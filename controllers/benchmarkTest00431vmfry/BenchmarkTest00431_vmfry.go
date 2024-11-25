//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: [89]
//-------------
//Semgrep analysis results: [89, 79, 319]
//Gosec analysis results: [676, 89, 703]
//CodeQL analysis results: [89]
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00431/BenchmarkTest00431.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/types/embedding.tmt with name embedding_anonymous_fields_positive 
//Used extensions: 
//Program:
package controllers

import (
	"database/sql"
	"fmt"
	"html"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00431 struct{}

func (b *BenchmarkTest00431) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00431")
	if param == "" {
		param = ""
	}

	bar := ""
	num := 106

	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

a := AnonFieldHolder{
    Value2: "const_str",
    Anon: Anon{
        Value1: bar,
    },
}
bar = a.Value1

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	results, err := db.Query(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	defer results.Close()

	w.Write([]byte("Your results are: <br>"))
	for results.Next() {
		var username string
		if err := results.Scan(&username); err != nil {
			http.Error(w, "Error processing query", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(encodeForHTML(username) + "<br>"))
	}
}

func encodeForHTML(s string) string {
	// Пример простой HTML-экранировки
	return html.EscapeString(s)
}

func main() {
	http.Handle("/sqli-00/BenchmarkTest00431", &BenchmarkTest00431{})
	http.ListenAndServe(":8080", nil)
}
