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

func init() {
	// initialization code here if necessary
}

type BenchmarkTest01107 struct{}

func (b *BenchmarkTest01107) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
	}
}

func (b *BenchmarkTest01107) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	for name, values := range r.Header {
		if shouldSkipHeader(name) {
			continue
		}
		if len(values) > 0 {
			param = name
			break
		}
	}

	bar := b.doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where id=%s", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Query execution error", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "JSON marshaling error", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func (b *BenchmarkTest01107) doSomething(param string) string {
	bar := ""
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func shouldSkipHeader(name string) bool {
	commonHeaders := []string{"User-Agent", "Accept", "Accept-Language", "Accept-Encoding"}
	for _, h := range commonHeaders {
		if h == name {
			return true
		}
	}
	return false
}
