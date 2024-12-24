package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02275 struct{}

func (b *BenchmarkTest02275) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02275")
	bar := doSomething(param)

tmpArrayUnique42 := []string{"", "", ""}
tmpArrayUnique42[0] = bar
ah := NewArrayHolderWithValues(tmpArrayUnique42)
bar = ah.Values[1]

	sqlStr := fmt.Sprintf("SELECT userid FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	var results int
	err := db.QueryRow(sqlStr).Scan(&results)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Your results are: %d", results)
}

func doSomething(param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
}
