package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

type BenchmarkTest02722Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02722Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02722Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest02722")

	bar := doSomething(param)

	c.Ctx.Request.Context().Value(bar) // Сохраняем значение в сессии

	output := fmt.Sprintf("Item: '%s' with value: 10340 saved in session.", bar)
	c.Ctx.ResponseWriter.Write([]byte(output))
}

func doSomething(param string) string {
	return param // Здесь можно использовать экранирование для HTML
}
