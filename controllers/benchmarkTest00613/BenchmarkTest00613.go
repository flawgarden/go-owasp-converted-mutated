package controllers

import (
	"database/sql"
	"encoding/json"
	"go-sec-code/models"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00613 struct{}

func (b *BenchmarkTest00613) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00613")
	if param == "" {
		param = "safe!"
	}

	bar := "safe!"
	map67704 := make(map[string]interface{})
	map67704["keyA-67704"] = "a_Value"
	map67704["keyB-67704"] = param
	map67704["keyC"] = "another_Value"

	bar = map67704["keyB-67704"].(string)
	bar = map67704["keyA-67704"].(string)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM user WHERE username=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	user := models.User{}
	err = stmt.QueryRow(bar).Scan(&user.Id, &user.Username, &user.Password)
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
