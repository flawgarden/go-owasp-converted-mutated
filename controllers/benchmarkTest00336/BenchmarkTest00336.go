package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest00336 struct {
	DB *sql.DB
}

func (b *BenchmarkTest00336) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest00336) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := r.Header["BenchmarkTest00336"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	var bar string
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	sqlStr := fmt.Sprintf("SELECT userid FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	var results int64
	err := b.DB.QueryRow(sqlStr).Scan(&results)

	if err == sql.ErrNoRows {
		http.Error(w, fmt.Sprintf("No results returned for query: %s", sqlStr), http.StatusNotFound)
	} else if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
	} else {
		output, _ := json.Marshal(map[string]interface{}{"results": results})
		w.Write(output)
	}
}

func NewBenchmarkTest00336(dataSourceName string) (*BenchmarkTest00336, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &BenchmarkTest00336{DB: db}, nil
}
