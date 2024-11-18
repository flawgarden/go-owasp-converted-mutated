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

type BenchmarkTest02258 struct {
	DB *sql.DB
}

func (b *BenchmarkTest02258) Init() {
	var err error
	b.DB, err = sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
}

func (b *BenchmarkTest02258) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	bar := doSomething(r, id)

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", bar)
	user := models.User{}
	err := b.DB.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
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

func doSomething(r *http.Request, param string) string {
	// Logic implementation using the param
	return param
}

func (b *BenchmarkTest02258) Close() {
	b.DB.Close()
}
