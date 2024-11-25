//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: [89]
//-------------
//Semgrep analysis results: [89]
//Gosec analysis results: [89, 703]
//CodeQL analysis results: [89]
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01222/BenchmarkTest01222.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/types/assertions.tmt with name type_assertion_in_switch_positive 
//Used extensions: 
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

type BenchmarkTest01222 struct{}

func (b *BenchmarkTest01222) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	param := ""
	if header := r.Header.Get("BenchmarkTest01222"); header != "" {
		param = header
	}
	param, _ = url.QueryUnescape(param)

var i interface{} = param
switch i.(type) {
case int:
    param = "PBWzW"
case string:
    param = param + "uwZxh"
default:
    param = "efoea"
}

	bar := new(Test).doSomething(r, param)

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

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map25738 := make(map[string]interface{})
	map25738["keyA-25738"] = "a-Value"
	map25738["keyB-25738"] = param
	map25738["keyC"] = "another-Value"
	bar = map25738["keyB-25738"].(string)
	return bar
}
