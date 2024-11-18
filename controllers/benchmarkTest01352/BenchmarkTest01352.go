package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01352Controller struct {
	web.Controller
}

func (c *BenchmarkTest01352Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01352Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01352")
	bar := c.doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.WriteString(bar)
}

func (c *BenchmarkTest01352Controller) doSomething(param string) string {
	return htmlEscape(param)
}

func htmlEscape(str string) string {
	return fmt.Sprintf("%s", str) // Имитация экранирования для примера
}
