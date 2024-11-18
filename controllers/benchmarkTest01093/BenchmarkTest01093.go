package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01093 struct{}

func (b *BenchmarkTest01093) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01093")
	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

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

	// Dummy response to mimic the original behavior
	w.Write([]byte("Query executed"))
}

func (b *BenchmarkTest01093) doSomething(param string) string {
	map18142 := map[string]interface{}{
		"keyA-18142": "a-Value",
		"keyB-18142": param,
		"keyC":       "another-Value",
	}
	return map18142["keyB-18142"].(string)
}
