package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01166 struct{}

func (b *BenchmarkTest01166) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.DoPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (b *BenchmarkTest01166) DoPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	if headers := r.Header["BenchmarkTest01166"]; len(headers) > 0 {
		param = headers[0]
	}
	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "JSON marshal error", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func (b *BenchmarkTest01166) doSomething(param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		sbxyz52006 := []rune(param)
		bar = string(append(sbxyz52006[:len(sbxyz52006)-1], 'Z'))
	}
	return bar
}
