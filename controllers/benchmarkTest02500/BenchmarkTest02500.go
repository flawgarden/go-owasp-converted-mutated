package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type BenchmarkTest02500Handler struct{}

func (h *BenchmarkTest02500Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02500")
	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", param)

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/goseccode")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	user := User{}
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

func main() {
	http.Handle("/weakrand-05/BenchmarkTest02500", &BenchmarkTest02500Handler{})
	http.ListenAndServe(":8080", nil)
}
