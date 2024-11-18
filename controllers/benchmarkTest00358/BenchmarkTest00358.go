package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00358 struct {
}

func (b *BenchmarkTest00358) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	id := r.URL.Query().Get("BenchmarkTest00358")
	if id == "" {
		id = ""
	}

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", id)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer db.Close()

	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Query execution error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "JSON marshaling error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Write(output)
}
