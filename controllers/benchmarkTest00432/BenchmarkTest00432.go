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

type BenchmarkTest00432 struct{}

func (b *BenchmarkTest00432) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00432")
	if param == "" {
		param = ""
	}

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			http.Error(w, "Error processing request", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	output, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error converting results to JSON", http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

func main() {
	http.Handle("/sqli-00/BenchmarkTest00432", &BenchmarkTest00432{})
	http.ListenAndServe(":8080", nil)
}
