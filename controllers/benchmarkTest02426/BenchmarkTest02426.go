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

type BenchmarkTest02426 struct {
	web.Controller
}

func (c *BenchmarkTest02426) Get() {
	c.Post()
}

func (c *BenchmarkTest02426) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")

	param := c.Ctx.Input.Query("BenchmarkTest02426")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id := bar // уязвимость SQL инъекции
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
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	bar := ""

	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}
