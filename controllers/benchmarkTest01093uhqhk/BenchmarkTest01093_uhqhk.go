// PASS
// Semgrep original results: [89]
// Gosec original results: [89]
// CodeQL original results: [89]
// Snyk original results: [89]
// -------------
// Semgrep analysis results: [89]
// Gosec analysis results: [89, 703]
// CodeQL analysis results: [89]
// Snyk analysis results: []
// Original file name: controllers/benchmarkTest01093/BenchmarkTest01093.go
// Original file CWE's: [89]
// Original file kind: fail
// Mutation info: Insert template from templates-db/languages/go/sensitivity/field/constructors.tmt with name class_with_array_initialization_neutral
// Used extensions: MACRO_Empty_string_Array -> []string{"", "", "", ""} | MACRO_Zero_Or_One -> 1
// Program:
package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01093 struct{}

func (b *BenchmarkTest01093) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01093")
	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(param)

	tmpArrayUnique42 := []string{"", "", "", ""}
	tmpArrayUnique42[0] = bar
	ah := NewArrayHolderWithValues(tmpArrayUnique42)
	bar = ah.Values[1]

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

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

	// Dummy response to mimic the original behavior
	w.Write([]byte("Query executed"))
}

func (b *BenchmarkTest01093) doSomething(param string) string {
	map18142 := map[string]interface{}{
		"keyA-18142": "a-Value",
		"keyB-18142": param,
		"keyC":       "another-Value",
	}
	return map18142["keyB-18142"].(string)
}
