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
	param := ""
	flag := true

	for name, values := range r.URL.Query() {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest00602" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}

	sqlStr := fmt.Sprintf("SELECT * FROM users WHERE username='foo' AND password='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var user models.User
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
