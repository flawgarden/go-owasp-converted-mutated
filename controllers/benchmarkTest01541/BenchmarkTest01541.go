package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest struct{}

func (b *BenchmarkTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	id := r.URL.Query().Get("BenchmarkTest01541")
	if id == "" {
		id = ""
	}

	bar := b.doSomething(id)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", bar)
	user := struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

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
	w.Write(output)
}

func (b *BenchmarkTest) doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}

func main() {
	http.Handle("/weakrand-03/BenchmarkTest01541", &BenchmarkTest{})
	http.ListenAndServe(":8080", nil)
}
