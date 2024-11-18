package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02063 struct {
	web.Controller
}

func (b *BenchmarkTest02063) Get() {
	b.Post()
}

func (b *BenchmarkTest02063) Post() {
	b.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := b.Ctx.Input.Header("BenchmarkTest02063")
	id := param

	sqlStr := fmt.Sprintf("select * from user where id=%s", id)
	user := models.User{}
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	b.Ctx.ResponseWriter.Write(output)
}
