package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

type BenchmarkTest01096 struct{}

func (b *BenchmarkTest01096) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var param string

	if r.Header.Get("BenchmarkTest01096") != "" {
		param = r.Header.Get("BenchmarkTest01096")
	}

	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(param)

	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", bar)

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/goseccode")
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
}

func (b *BenchmarkTest01096) doSomething(param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C':
	case 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	return bar
}
