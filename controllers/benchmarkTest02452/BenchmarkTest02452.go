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

type SqlInjectionVuln2Controller struct{}

func (c *SqlInjectionVuln2Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest02452")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			http.Error(w, "Error scanning rows", http.StatusInternalServerError)
			return
		}
		results = append(results, user)
	}

	output, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func doSomething(param string) string {
	a91263 := param
	b91263 := a91263 + " SafeStuff"
	return b91263[:len(b91263)-1]
}
