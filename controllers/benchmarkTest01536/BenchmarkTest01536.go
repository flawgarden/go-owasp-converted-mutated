package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01536 struct{}

func (bt *BenchmarkTest01536) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("BenchmarkTest01536")
	if id == "" {
		id = "0"
	}

	user := getUserByID(id)

	output, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func getUserByID(id string) User {
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", id) // Vulnerable to SQL Injection
	var user User
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}
	return user
}

type User struct {
	Id       int
	Username string
	Password string
}

func main() {
	http.Handle("/weakrand-03/BenchmarkTest01536", &BenchmarkTest01536{})
	http.ListenAndServe(":8080", nil)
}
