package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02248 struct {
	db *sql.DB
}

func NewBenchmarkTest02248() (*BenchmarkTest02248, error) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}
	return &BenchmarkTest02248{db: db}, nil
}

func (b *BenchmarkTest02248) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (b *BenchmarkTest02248) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest02248")

	bar := b.doSomething(param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   true,
		HttpOnly: true,
		Path:     r.URL.Path,
	}
	http.SetCookie(w, &cookie)

	fmt.Fprintf(w, "Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: true", bar)
}

func (b *BenchmarkTest02248) doSomething(param string) string {
	if param != "" {
		return param
	}
	return ""
}

func main() {
	handler, err := NewBenchmarkTest02248()
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8080", handler)
}
