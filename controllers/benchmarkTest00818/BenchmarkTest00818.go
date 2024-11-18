package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00818 struct{}

func (bt *BenchmarkTest00818) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest00818="
	paramLoc := strings.Index(queryString, paramval)

	if paramLoc == -1 {
		w.Write([]byte(`getQueryString() couldn't find expected parameter 'BenchmarkTest00818' in query string.`))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(param, "&")
	if ampersandLoc != -1 {
		param = param[:ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := param
	if param != "" && len(param) > 1 {
		bar = param[:len(param)-1]
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE username='%s'", bar)
	var user models.User
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
