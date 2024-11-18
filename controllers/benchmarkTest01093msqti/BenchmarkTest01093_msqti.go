//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: [89]
//-------------
//Semgrep analysis results: [89]
//Gosec analysis results: [89, 703]
//CodeQL analysis results: [89]
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01093/BenchmarkTest01093.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/varargs/varargs.tmt with name varargs_get_string_with_index_negative 
//Used extensions: 
//Program:
package controllers

import (
"database/sql"
"fmt"
"net/http"
"net/url"
_ "github.com/go-sql-driver/mysql"
"strings"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01093 struct{}

func (b *BenchmarkTest01093) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01093")
	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)

sqlStr = getStringWithIndex(1, sqlStr, "QmiTY")

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

func getFirstString(lines ...string) string {
    return getStringWithIndex(0, lines...)
}

func getStringWithIndex(ind int, lines ...string) string {
    return lines[ind]
}

func getFirstStringFromArray(lines ...string) string {
    return lines[0]
}

func varargsWithGenerics[T any](elements ...T) T {
    return elements[0]
}

func combineStrings(strs ...string) string {
    return strings.Join(strs, ", ")
}


