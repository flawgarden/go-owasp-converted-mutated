// FAIL
// Semgrep original results: [89]
// Gosec original results: [89]
// CodeQL original results: [89]
// Snyk original results: [89]
// -------------
// Semgrep analysis results: [89, 79, 319]
// Gosec analysis results: [676, 89, 703]
// CodeQL analysis results: [89]
// Snyk analysis results: []
// Original file name: controllers/benchmarkTest00516/BenchmarkTest00516.go
// Original file CWE's: [89]
// Original file kind: fail
// Mutation info: Insert template from templates-db/languages/go/sensitivity/defer.tmt with name simple_defer_neutral
// Used extensions:
// Program:
package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00516 struct{}

func (b *BenchmarkTest00516) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00516")
	bar := ""
	guess := "ABC"
	switchTarget := guess[2]

	defer func() {
		param = "MsLmP"
	}()

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []map[string]interface{}{}
	for rows.Next() {
		var id int
		var username, password string
		err = rows.Scan(&id, &username, &password)
		if err != nil {
			http.Error(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		users = append(users, map[string]interface{}{
			"id":       id,
			"username": username,
			"password": password,
		})
	}

	output, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func main() {
	http.Handle("/sqli-01/BenchmarkTest00516", &BenchmarkTest00516{})
	http.ListenAndServe(":8080", nil)
}

func foo(f string) (s string) {
	defer func() {
		s = "constant_string"
	}()
	s = f + " suffix"
	return s
}

func foo2(f string) (s string) {
	defer func() {
		s = s + f
	}()
	s = f + " suffix"
	return s
}
