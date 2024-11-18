//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: [89]
//-------------
//Semgrep analysis results: [89, 79]
//Gosec analysis results: [89, 703]
//CodeQL analysis results: [89]
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00678/BenchmarkTest00678.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/types/assertions.tmt with name type_assertion_in_switch_negative 
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

type BenchmarkTest00678 struct{}

func (b *BenchmarkTest00678) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00678")

var i interface{} = 174807775
switch i.(type) {
case int:
    param = "PaluX"
case string:
    param = param + "PXIQv"
default:
    param = "TnpoQ"
}

	if param == "" {
		param = ""
	}

	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	sqlStr := fmt.Sprintf("SELECT userid FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var results int64
	err = db.QueryRow(sqlStr).Scan(&results)
	if err != nil {
		http.Error(w, "No results returned for query: "+sqlStr, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"results": results}
	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
