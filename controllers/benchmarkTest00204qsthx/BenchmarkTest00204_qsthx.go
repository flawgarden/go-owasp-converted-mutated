package controllers

import (
"database/sql"
"fmt"
"net/http"
"net/url"
_ "github.com/go-sql-driver/mysql"
"sync"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00204 struct{}

func (b *BenchmarkTest00204) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.Header.Get("BenchmarkTest00204")

message123 := make(chan string, 1)
message123 <- param

var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    rmsg := <-message123
    message123 <- rmsg + "constant_string"
}()

wg.Wait()

param = <-message123

	param, _ = url.QueryUnescape(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	bar := param
	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo', '%s')", bar)

	_, err = db.Exec(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Update complete"))
}
