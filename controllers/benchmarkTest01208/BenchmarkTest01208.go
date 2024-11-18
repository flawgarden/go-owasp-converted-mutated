package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01208 struct {
	DB *sql.DB
}

func (b *BenchmarkTest01208) Init() {
	var err error
	b.DB, err = sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
}

func (b *BenchmarkTest01208) DoGet(w http.ResponseWriter, r *http.Request) {
	b.DoPost(w, r)
}

func (b *BenchmarkTest01208) DoPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if h := r.Header.Get("BenchmarkTest01208"); h != "" {
		param = h
	}

	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(r, param)

	sqlStr := fmt.Sprintf("{call %s}", bar)

	var user models.User
	err := b.DB.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func (b *BenchmarkTest01208) doSomething(r *http.Request, param string) string {
	bar := ""
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}
	return bar
}
