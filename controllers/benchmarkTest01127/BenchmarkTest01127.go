package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest struct {
	db *sql.DB
}

func NewBenchmarkTest() *BenchmarkTest {
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	return &BenchmarkTest{db: db}
}

func (bt *BenchmarkTest) Close() {
	bt.db.Close()
}

func (bt *BenchmarkTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	sqlStr := fmt.Sprintf("select * from user where id=%s", id)
	user := models.User{}
	err := bt.db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
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
