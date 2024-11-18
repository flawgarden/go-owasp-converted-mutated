package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00185 struct{}

func (b *BenchmarkTest00185) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest00185")
	param, _ = url.QueryUnescape(param)

	bar := "safe!"
	dataMap := make(map[string]interface{})
	dataMap["keyA-60659"] = "a_Value"
	dataMap["keyB-60659"] = param
	dataMap["keyC"] = "another_Value"
	bar = dataMap["keyB-60659"].(string)
	bar = dataMap["keyA-60659"].(string)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	id, err := strconv.Atoi(bar) // assumed bar would hold a numeric value for SQL query
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%d", id)
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
	http.Handle("/weakrand-00/BenchmarkTest00185", &BenchmarkTest00185{})
	http.ListenAndServe(":8080", nil)
}
