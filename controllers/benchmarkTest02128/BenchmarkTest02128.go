package controllers

import (
	"database/sql"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02128Controller struct {
	web.Controller
}

func (c *BenchmarkTest02128Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest02128Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest02128Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02128")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	c.Ctx.Output.Body([]byte(bar))
}

func doSomething(param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}

func init() {
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		panic(err)
	}
}
