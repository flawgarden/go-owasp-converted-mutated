package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01354Controller struct {
	web.Controller
}

func (c *BenchmarkTest01354Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest01354Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest01354Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01354")
	bar := c.doSomething(param)

	id := bar
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where id=%s", id)
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

func (c *BenchmarkTest01354Controller) doSomething(param string) string {
	bar := ""
	if param != "" {
		bar = param
	}

	return bar
}
