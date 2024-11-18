package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02434 struct{}

func (b *BenchmarkTest02434) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	id := r.URL.Query().Get("BenchmarkTest02434")
	if id == "" {
		id = ""
	}

	bar := doSomething(r, id)

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

func doSomething(r *http.Request, param string) string {
	bar := "safe!"
	mapData := map[string]interface{}{
		"keyA": "a-Value",
		"keyB": param,
		"keyC": "another-Value",
	}
	bar = mapData["keyB"].(string)
	return bar
}

func main() {
	http.Handle("/weakrand-05/BenchmarkTest02434", &BenchmarkTest02434{})
	http.ListenAndServe(":8080", nil)
}
