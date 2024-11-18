package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type BenchmarkTest01475 struct{}

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func (bt *BenchmarkTest01475) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := ""
	flag := true

	for name, values := range r.Form {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest01475" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := bt.doSomething(r, param)
	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo', '%s')", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if _, err := db.Exec(sqlStr); err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Update complete"))
}

func (bt *BenchmarkTest01475) doSomething(r *http.Request, param string) string {
	// Dummy implementation for the purpose of the example
	return param
}
