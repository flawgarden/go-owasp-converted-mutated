//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
// Semgrep original results: [89]
// Gosec original results: [89]
// CodeQL original results: [89]
// Snyk original results: [89]
// -------------
// Semgrep analysis results: [89, 79]
// Gosec analysis results: [89, 703]
// CodeQL analysis results: [89]
// Snyk analysis results: []
// Original file name: controllers/benchmarkTest00850/BenchmarkTest00850.go
// Original file CWE's: [89]
// Original file kind: fail
// Mutation info: Insert template from templates-db/languages/go/sensitivity/types/assertions.tmt with name type_assertion_with_struct_pointer_positive
// Used extensions:
// Program:
package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00850 struct {
}

func (b *BenchmarkTest00850) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest00850="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest00850"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc+paramLoc]
	}
	param, err := url.QueryUnescape(param)
	if err != nil {
		http.Error(w, "Failed to decode parameter", http.StatusBadRequest)
		return
	}

	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:]
		bar = valuesList[0]

		var i123 interface{} = &EmbeddedStruct{Field1: bar}
		if ptr, ok := i123.(*EmbeddedStruct); ok {
			bar = ptr.Field1
		} else {
			bar = "ZwafJ"
		}

	}

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var user models.User
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
