package controllers

import (
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

type BenchmarkTest02237Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02237Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02237Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Query("BenchmarkTest02237")
	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.WriteString(bar)
}

func doSomething(param string) string {
	// Эмуляция ESAPI encoder для безопасной передачи
	return param // Вернуть параметр без изменений, вместо кодирования для упрощения
}
