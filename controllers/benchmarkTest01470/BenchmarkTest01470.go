package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01470 struct{}

func (b *BenchmarkTest01470) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := ""
	flag := true
	for name, values := range r.Form {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest01470" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := b.doSomething(param)

	sqlStr := fmt.Sprintf("SELECT TOP 1 userid FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	result, err := db.Query(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	defer result.Close()

	var results []models.User
	for result.Next() {
		var user models.User
		if err := result.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			http.Error(w, "Error scanning results", http.StatusInternalServerError)
			return
		}
		results = append(results, user)
	}

	output, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error marshaling results", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (b *BenchmarkTest01470) doSomething(param string) string {
	return param
}
