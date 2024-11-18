package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02503 struct {
	http.Handler
}

func (bt *BenchmarkTest02503) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("BenchmarkTest02503")
	bar := doSomething(id)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where id='%s'", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
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

func doSomething(param string) string {
	a := param
	b := a + " SafeStuff"
	b = b[:len(b)-1] + "Chars"
	return b // Simplified for this example
}
