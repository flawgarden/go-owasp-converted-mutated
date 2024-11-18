package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01149 struct{}

func (b *BenchmarkTest01149) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := ""

	headers := r.Header["BenchmarkTest01149"]
	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", bar)
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
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		return param
	case 'B':
		return "bobs_your_uncle"
	case 'C', 'D':
		return param
	default:
		return "bobs_your_uncle"
	}
}

func main() {
	http.Handle("/crypto-01/BenchmarkTest01149", &BenchmarkTest01149{})
	http.ListenAndServe(":8080", nil)
}
