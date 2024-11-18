package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/url"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type SqlInjectionVuln2Controller struct{}

func (c *SqlInjectionVuln2Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.Header.Get("BenchmarkTest01095")
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	sqlStr := "SELECT * from USERS where USERNAME='foo' and PASSWORD='" + bar + "'"
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

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Username, &user.Password); err != nil {
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
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := ""
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
