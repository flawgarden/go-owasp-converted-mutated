package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"
	"net/url"
	"strings"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01648 struct{}

func (bt *BenchmarkTest01648) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest01648="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest01648' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var user models.User
	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", bar)
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	w.Write(output)
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	return htmlEscape(param)
}

func htmlEscape(s string) string {
	return strings.ReplaceAll(s, "<", "&lt;")
}
