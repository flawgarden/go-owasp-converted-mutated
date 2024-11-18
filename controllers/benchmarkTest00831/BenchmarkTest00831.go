package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00831 struct{}

func (b *BenchmarkTest00831) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := r.URL.Query()
	param := queryString.Get("BenchmarkTest00831")
	if param == "" {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest00831' in query string.", http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error connecting to the database.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", param)
	var user models.User
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Error querying the database.", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error marshaling user data.", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}
