package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01548 struct{}

func (b *BenchmarkTest01548) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01548")
	if param == "" {
		param = ""
	}
	bar := b.doSomething(r, param)

	r.Context().Value("session").(map[string]interface{})["userid"] = bar

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprintf(w, "Item: 'userid' with value: '%s' saved in session.", htmlEscape(bar))
}

func (b *BenchmarkTest01548) doSomething(r *http.Request, param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	return bar
}

func htmlEscape(s string) string {
	output, _ := json.Marshal(s)
	return string(output)
}

func initDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
