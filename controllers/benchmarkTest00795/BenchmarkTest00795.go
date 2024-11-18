package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func BenchmarkTest00795(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	query := r.URL.Query()
	paramval := "BenchmarkTest00795="
	param := query.Get("BenchmarkTest00795")

	if param == "" {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", paramval), http.StatusBadRequest)
		return
	}

	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}

	id, err := strconv.Atoi(bar)
	if err != nil {
		http.Error(w, "Invalid ID format.", http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Database connection error.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var user User
	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%d", id)
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "User not found.", http.StatusNotFound)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error marshaling response.", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func main() {
	http.HandleFunc("/hash-00/BenchmarkTest00795", BenchmarkTest00795)
	http.ListenAndServe(":8080", nil)
}
