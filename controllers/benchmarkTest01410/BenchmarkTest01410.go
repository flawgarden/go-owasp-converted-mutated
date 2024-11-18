package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest struct {
}

func (b *BenchmarkTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var param string
	flag := true

	for name, values := range r.Form {
		if values != nil {
			for _, value := range values {
				if value == "BenchmarkTest01410" {
					param = name
					flag = false
					break
				}
			}
		}
		if !flag {
			break
		}
	}

	bar := b.doSomething(param)

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

func (b *BenchmarkTest) doSomething(param string) string {
	// Simple ? condition that assigns param to bar on false condition
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
}
