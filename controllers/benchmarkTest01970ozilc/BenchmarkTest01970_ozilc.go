//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: [89]
//-------------
//Semgrep analysis results: [319, 89, 79]
//Gosec analysis results: [676, 89, 703]
//CodeQL analysis results: [89, 116, 79]
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01970/BenchmarkTest01970.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/constructors.tmt with name class_with_array_initialization_positive 
//Used extensions: MACRO_Empty_string_Array -> []string{"", ""}
//Program:
package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func main() {
	http.HandleFunc("/sqli-04/BenchmarkTest01970", BenchmarkTest01970)
	http.ListenAndServe(":8080", nil)
}

func BenchmarkTest01970(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01970")
	param, _ = url.QueryUnescape(param)

	bar := doSomething(r, param)

tmpArrayUnique42 := []string{"", ""}
tmpArrayUnique42[0] = bar
ah := NewArrayHolderWithValues(tmpArrayUnique42)
bar = ah.Values[0]

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

	fmt.Fprintf(w, "Executed SQL: %s", sqlStr)
}

func doSomething(r *http.Request, param string) string {
	bar := ""
	guess := "ABC"
	switchTarget := guess[2]

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

	return bar
}
