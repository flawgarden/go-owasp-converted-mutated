package controllers

import (
	"database/sql"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01513Controller struct {
	web.Controller
}

func (c *BenchmarkTest01513Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01513Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest01513")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(param)
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.Output.Body([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := ""

	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func init() {
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
