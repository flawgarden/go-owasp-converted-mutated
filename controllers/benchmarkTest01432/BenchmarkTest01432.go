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

type BenchmarkTest01432 struct{}

func (b *BenchmarkTest01432) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	param := ""
	flag := true
	for name, values := range r.Form {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest01432" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := b.doSomething(param)

	id := bar
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", id)
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

func (b *BenchmarkTest01432) doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		sbxyz18866 := strings.Builder{}
		sbxyz18866.WriteString(param)
		bar = sbxyz18866.String()[:len(param)-1] + "Z"
	}
	return bar
}
