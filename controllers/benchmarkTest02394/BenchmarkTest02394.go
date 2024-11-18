package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02394 struct {
	web.Controller
}

func (c *BenchmarkTest02394) Get() {
	c.Post()
}

func (c *BenchmarkTest02394) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02394")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	_, _ = fmt.Fprintf(c.Ctx.ResponseWriter, bar)
}

func doSomething(param string) string {
	return param // Здесь следует применять экранирование, если это необходимо
}

func main() {
	web.Router("/xss-04/BenchmarkTest02394", &BenchmarkTest02394{})
	web.Run() // Запуск сервера
}
