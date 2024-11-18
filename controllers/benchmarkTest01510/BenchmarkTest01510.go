package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func main() {
	http.HandleFunc("/xss-03/BenchmarkTest01510", BenchmarkTest01510)
	http.ListenAndServe(":8080", nil)
}

func BenchmarkTest01510(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	id := r.URL.Query().Get("BenchmarkTest01510")
	if id == "" {
		id = ""
	}

	bar := doSomething(r, id)

	w.Header().Set("X-XSS-Protection", "0")
	fmt.Fprintf(w, "Formatted like: %s and %s.", bar, "b")
}

func doSomething(r *http.Request, param string) string {
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var user models.User
	sqlStr := fmt.Sprintf("select * from user where id=%s", param)
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	return string(output)
}
