package controllers

import (
	"database/sql"
	"fmt"
	"html"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02576 struct {
	DB *sql.DB
}

func NewBenchmarkTest02576() (*BenchmarkTest02576, error) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}
	return &BenchmarkTest02576{DB: db}, nil
}

func (b *BenchmarkTest02576) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02576")

	bar := doSomething(param)

	hashedValue := hashValue(bar)

	fileTarget := "passwordFile.txt"
	file, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Could not open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("hash_value=%s\n", hashedValue)); err != nil {
		http.Error(w, "Could not write to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprintf(w, "Sensitive value '%s' hashed and stored<br/>", html.EscapeString(bar))
}

func doSomething(param string) string {
	bar := ""
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}

func hashValue(input string) string {
	// Implement your hashing logic here
	return input // Placeholder for the actual hashing logic
}
