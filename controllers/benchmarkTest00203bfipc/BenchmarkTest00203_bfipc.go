//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: [89]
//-------------
//Semgrep analysis results: [89, 319]
//Gosec analysis results: [676, 89]
//CodeQL analysis results: [89]
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00203/BenchmarkTest00203.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/concurrency/concurrency.tmt with name thread_set_positive 
//Used extensions: 
//Program:
package controllers

import (
"database/sql"
"fmt"
"net/http"
"net/url"
"os"
_ "github.com/go-sql-driver/mysql"
"sync"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00203 struct{}

func (b *BenchmarkTest00203) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Del("Content-Length")
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("BenchmarkTest00203") != "" {
		param = r.Header.Get("BenchmarkTest00203")

w := &Wrapper[string]{Value: param}
task := NewSettingTask(w, param)
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    task.Run()
}()
wg.Wait()
param = w.Value

	}
	param, _ = url.QueryUnescape(param)

	var bar string
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

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

	fmt.Fprintln(w, "Update complete!")
}

func main() {
	http.Handle("/sqli-00/BenchmarkTest00203", &BenchmarkTest00203{})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start: ", err)
		os.Exit(1)
	}
}
