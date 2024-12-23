package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02362 struct{}

func (b *BenchmarkTest02362) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var param string
	flag := true

	for name, value := range r.URL.Query() {
		if flag {
			for _, v := range value {
				if v == "BenchmarkTest02362" {
					param = name
					flag = false
				}
			}
		}
	}

	bar := doSomething(r, param)

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)

tmpUnique42 := ""
switch "BWhYH" {
case "NxNFX":
    sqlStr = ""
default:
    sqlStr = tmpUnique42
}

	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	results, err := db.Query(sqlStr)
	if err != nil {
		http.Error(w, "Query execution error", http.StatusInternalServerError)
		return
	}
	defer results.Close()

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprintln(w, "Your results are: ")
	for results.Next() {
		var user models.User
		if err := results.Scan(&user.Username); err != nil {
			http.Error(w, "Error scanning results", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%s ", user.Username)
	}
}

func doSomething(r *http.Request, param string) string {
	bar := ""
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}
	return bar
}
