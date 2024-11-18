package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01555 struct{}

func (bt *BenchmarkTest01555) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	id := r.URL.Query().Get("BenchmarkTest01555")
	if id == "" {
		id = ""
	}

	bar := doSomething(r, id)

	sqlStr := fmt.Sprintf("SELECT userid FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var results int64
	err = db.QueryRow(sqlStr).Scan(&results)
	if err != nil {
		http.Error(w, "No results returned for query.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Your results are: %d", results)
}

func doSomething(r *http.Request, param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}
	return bar
}

func main() {
	http.Handle("/sqli-03/BenchmarkTest01555", &BenchmarkTest01555{})
	http.ListenAndServe(":8080", nil)
}
