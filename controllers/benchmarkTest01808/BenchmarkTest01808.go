package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01808 struct{}

func (b *BenchmarkTest01808) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/sqli-03/BenchmarkTest01808", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	param := r.FormValue("BenchmarkTest01808")
	bar := b.doSomething(param)

	sqlStr := fmt.Sprintf("SELECT userid from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var results int
	err = db.QueryRow(sqlStr).Scan(&results)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	_, _ = fmt.Fprintf(w, "Your results are: %d", results)
}

func (b *BenchmarkTest01808) doSomething(param string) string {
	bar := "bob's your uncle"
	guess := "ABC"
	switchTarget := guess[1]

	switch switchTarget {
	case 'A':
		bar = param
	case 'C', 'D':
		bar = param
	case 'B':
		bar = "bob"
	}
	return bar
}

func main() {
	http.Handle("/sqli-03/BenchmarkTest01808", &BenchmarkTest01808{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
