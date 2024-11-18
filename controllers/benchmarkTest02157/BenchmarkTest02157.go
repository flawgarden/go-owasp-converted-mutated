package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"
	"strconv"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02157 struct {
	DB *sql.DB
}

func NewBenchmarkTest02157() (*BenchmarkTest02157, error) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}
	return &BenchmarkTest02157{DB: db}, nil
}

func (b *BenchmarkTest02157) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("BenchmarkTest02157")
	if param == "" {
		param = "0"
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%d", id)
	user := models.User{}
	err = b.DB.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
