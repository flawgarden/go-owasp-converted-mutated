package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

type BenchmarkTest00205 struct {
	DB *sql.DB
}

func (b *BenchmarkTest00205) Init() {
	var err error
	b.DB, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/goseccode")
	if err != nil {
		panic(err)
	}
}

func (b *BenchmarkTest00205) DoGet(w http.ResponseWriter, r *http.Request) {
	b.DoPost(w, r)
}

func (b *BenchmarkTest00205) DoPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	param := r.Header.Get("BenchmarkTest00205")
	param, _ = url.QueryUnescape(param)

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}

	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", bar)

	_, err := b.DB.Exec(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Update complete"}
	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}
