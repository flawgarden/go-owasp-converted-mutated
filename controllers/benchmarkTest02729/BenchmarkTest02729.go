package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02729 struct{}

func (b *BenchmarkTest02729) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02729")
	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		http.Error(w, "Error preparing statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec("foo")
	if err != nil {
		http.Error(w, "Error executing statement", http.StatusInternalServerError)
		return
	}

	// Handle results here
}

func doSomething(param string) string {
	bar := "safe!"
	map87005 := make(map[string]interface{})
	map87005["keyA-87005"] = "a-Value"
	map87005["keyB-87005"] = param
	map87005["keyC"] = "another-Value"
	bar = map87005["keyB-87005"].(string)

	return bar
}

func main() {
	http.Handle("/sqli-06/BenchmarkTest02729", &BenchmarkTest02729{})
	http.ListenAndServe(":8080", nil)
}
