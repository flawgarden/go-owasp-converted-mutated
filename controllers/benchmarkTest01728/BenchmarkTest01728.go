package controllers

import (
	"database/sql"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

type BenchmarkTest01728 struct {
	DB *sql.DB
}

func NewBenchmarkTest01728(dataSourceName string) (*BenchmarkTest01728, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &BenchmarkTest01728{DB: db}, nil
}

func (b *BenchmarkTest01728) DoGet(w http.ResponseWriter, r *http.Request) {
	b.DoPost(w, r)
}

func (b *BenchmarkTest01728) DoPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest01728="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramval)
		if paramLoc < 0 {
			http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest01728'", http.StatusBadRequest)
			return
		}
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := len(queryString)
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := new(Test).DoSomething(param)

	sql := "SELECT * from USERS where USERNAME='foo' and PASSWORD='" + bar + "'"

	_, err := b.DB.Exec(sql)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
}

type Test struct{}

func (t *Test) DoSomething(param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
}
