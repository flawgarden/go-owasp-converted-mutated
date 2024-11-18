package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01662 struct {
	db *sql.DB
}

func NewBenchmarkTest01662() (*BenchmarkTest01662, error) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}
	return &BenchmarkTest01662{db: db}, nil
}

func (b *BenchmarkTest01662) Get(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest01662="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramval)
		if paramLoc < 0 || queryString[paramLoc:len(paramval)+paramLoc] != paramval {
			http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest01662' in query string.", http.StatusBadRequest)
			return
		}
	}

	param := queryString[paramLoc+len(paramval):]
	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	fmt.Fprintf(w, "Formatted like: %s and %s.", bar, "b")
}

func (b *BenchmarkTest01662) doSomething(param string) string {
	var bar string
	if param != "" {
		bar = string(param) // Пример обработки параметра, можно заменить на другую логику.
	}
	return bar
}
