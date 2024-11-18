package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01462 struct {
	web.Controller
}

func (b *BenchmarkTest01462) Get() {
	b.doPost()
}

func (b *BenchmarkTest01462) Post() {
	b.doPost()
}

func (b *BenchmarkTest01462) doPost() {
	b.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := b.Ctx.Request.URL.Query()

	for name, values := range names {
		if !flag {
			break
		}
		for _, value := range values {
			if value == "BenchmarkTest01462" {
				param = name
				flag = false
				break
			}
		}
	}

	bar := new(Test).doSomething(param)

	sqlStr := "SELECT * FROM user WHERE id = ?"
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(b.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	user := models.User{}
	err = db.QueryRow(sqlStr, bar).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(b.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(b.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	b.Ctx.ResponseWriter.Write(output)
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := param
	return bar
}
