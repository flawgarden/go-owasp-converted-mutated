package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02018 struct {
}

func (b *BenchmarkTest02018) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.Header.Get("BenchmarkTest02018")

	param, err := url.QueryUnescape(param)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}


	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", param)
	user := &models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error processing JSON", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func doSomething(param string) string {
	a50268 := param
	b50268 := a50268 + " SafeStuff"
	b50268 = b50268[:len(b50268)-1] + "Chars"
	return b50268
}
