package controllers

import (
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

type BenchmarkTest01420Controller struct {
	web.Controller
}

func (c *BenchmarkTest01420Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01420Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := c.Ctx.Request.URL.Query()
	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest01420" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := c.doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func (c *BenchmarkTest01420Controller) doSomething(param string) string {
	// Эмуляция безопасного ввода
	return param // Необходима замена на правильную обработку
}
