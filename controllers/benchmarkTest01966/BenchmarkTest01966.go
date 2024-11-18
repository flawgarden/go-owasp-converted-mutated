package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01966 struct{}

func (b *BenchmarkTest01966) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	responseType := "text/html;charset=UTF-8"
	w.Header().Set("Content-Type", responseType)

	param := ""
	if r.Header.Get("BenchmarkTest01966") != "" {
		param = r.Header.Get("BenchmarkTest01966")
	}

	param, _ = url.QueryUnescape(param)
	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT TOP 1 USERNAME from USERS where USERNAME='foo' and PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error connecting to the database.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var result string
	err = db.QueryRow(sqlStr).Scan(&result)
	if err != nil {
		http.Error(w, "No results returned for query: "+sqlStr, http.StatusNotFound)
		return
	}

	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

func doSomething(param string) string {
	bar := ""

	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}

func main() {
	http.Handle("/sqli-04/BenchmarkTest01966", &BenchmarkTest01966{})
	http.ListenAndServe(":8080", nil)
}
