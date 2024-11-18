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

type BenchmarkTest02148 struct{}

func (b *BenchmarkTest02148) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	id := r.URL.Query().Get("BenchmarkTest02148")
	if id == "" {
		id = "0"
	}

	bar := b.doSomething(id)

	sqlStr := fmt.Sprintf("select * from user where id=%s", bar)
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
	w.Write(output)
}

func (b *BenchmarkTest02148) doSomething(param string) string {
	a17988 := param
	b17988 := a17988 + " SafeStuff"
	b17988 = b17988[:len(b17988)-5] + "Chars"
	return b17988
}
