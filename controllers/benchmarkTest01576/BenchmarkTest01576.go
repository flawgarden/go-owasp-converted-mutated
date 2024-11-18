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

type BenchmarkTest01576 struct{}

func (b *BenchmarkTest01576) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("BenchmarkTest01576")
	bar := new(Test).doSomething(id)

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	user := models.User{}
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

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := param
	return bar
}

func main() {
	http.Handle("/hash-01/BenchmarkTest01576", &BenchmarkTest01576{})
	http.ListenAndServe(":8080", nil)
}
