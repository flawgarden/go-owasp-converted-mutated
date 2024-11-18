package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00540 struct{}

func (bt *BenchmarkTest00540) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	id := r.URL.Query().Get("id")
	var param string
	if id == "BenchmarkTest00540" {
		param = id
	}

	bar := param

	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", bar)
	user := struct {
		Id       int
		Username string
		Password string
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

func main() {
	http.Handle("/hash-00/BenchmarkTest00540", &BenchmarkTest00540{})
	http.ListenAndServe(":8080", nil)
}
