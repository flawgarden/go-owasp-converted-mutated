package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02273 struct{}

func (b *BenchmarkTest02273) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	param := r.FormValue("BenchmarkTest02273")

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)
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

	_, _ = w.Write([]byte("No results can be displayed for query: " + sqlStr))
}

func doSomething(param string) string {
	guess := "ABC"
	switchTarget := guess[2]

	var bar string
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

func main() {
	http.Handle("/sqli-04/BenchmarkTest02273", &BenchmarkTest02273{})
	http.ListenAndServe(":8080", nil)
}
