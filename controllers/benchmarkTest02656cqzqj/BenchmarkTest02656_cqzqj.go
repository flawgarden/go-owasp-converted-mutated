//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: [89]
//-------------
//Semgrep analysis results: [319, 89]
//Gosec analysis results: [676, 89, 703]
//CodeQL analysis results: [89]
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest02656/BenchmarkTest02656.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/varargs/varargs.tmt with name varargs_get_string_with_index_positive 
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

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("/sqli-06/BenchmarkTest02656", BenchmarkTest02656)
	http.ListenAndServe(":8080", nil)
}

func BenchmarkTest02656(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02656="

queryString = getStringWithIndex(0, queryString, "xHvEE")

	paramLoc := -1

	if queryString != "" {
		paramLoc = len(queryString) - len(paramval)
	}

	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02656' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := indexOfAmpersand(queryString, paramLoc)
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)
	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", bar)

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

	w.Write([]byte("Update complete"))
}

func indexOfAmpersand(queryString string, start int) int {
	for i := start; i < len(queryString); i++ {
		if queryString[i] == '&' {
			return i
		}
	}
	return -1
}

func doSomething(param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
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


