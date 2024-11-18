package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01969 struct{}

func (bt *BenchmarkTest01969) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	if r.Header.Get("BenchmarkTest01969") != "" {
		param = r.Header.Get("BenchmarkTest01969")
	}

	decodedParam, _ := url.QueryUnescape(param)
	bar := doSomething(r, decodedParam)

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	output, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func doSomething(r *http.Request, param string) string {
	bar := ""
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
