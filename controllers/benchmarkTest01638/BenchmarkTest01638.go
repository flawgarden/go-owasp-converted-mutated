package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01638 struct {
}

func (b *BenchmarkTest01638) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest01638="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest01638' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc+paramLoc]
	}
	param = strings.TrimSpace(param)

	bar := new(Test).doSomething(r, param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}
	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

type Test struct {
}

func (t *Test) doSomething(r *http.Request, param string) string {
	a61263 := param
	b61263 := strings.Builder{}
	b61263.WriteString(a61263)
	b61263.WriteString(" SafeStuff")
	b61263Str := b61263.String()[:len(b61263.String())-1]
	return b61263Str // пример использования для дальнейшей обработки
}
