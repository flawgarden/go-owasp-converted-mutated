package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"
)

type BenchmarkTest02295Controller struct {
}

func (c *BenchmarkTest02295Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.doPost(w, r)
	} else if r.Method == http.MethodPost {
		c.doPost(w, r)
	}
}

func (c *BenchmarkTest02295Controller) doPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := ""

	for name, values := range r.Form {
		for _, value := range values {
			if value == "BenchmarkTest02295" {
				param = name
			}
		}
	}

	bar := doSomething(param)

	source := "root:password@tcp(127.0.0.1:3306)/goseccode"
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
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
