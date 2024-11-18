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

type BenchmarkTest01479 struct {
	db *sql.DB
}

func NewBenchmarkTest() (*BenchmarkTest01479, error) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}
	return &BenchmarkTest01479{db: db}, nil
}

func (b *BenchmarkTest01479) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var param string
	for name, values := range r.Form {
		for _, value := range values {
			if value == "BenchmarkTest01479" {
				param = name
				break
			}
		}
	}

	bar := b.doSomething(param)

	defer b.db.Close()
	sqlStr := fmt.Sprintf("select * from user where id='%s'", bar)
	user := models.User{}
	err := b.db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
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

func (b *BenchmarkTest01479) doSomething(param string) string {
	bar := ""
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
