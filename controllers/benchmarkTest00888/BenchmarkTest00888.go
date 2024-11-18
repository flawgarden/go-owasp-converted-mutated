package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00888Controller struct {
	web.Controller
}

func (c *BenchmarkTest00888Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00888Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00888")
	bar := sanitizeInput(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func sanitizeInput(param string) string {
	// Простейшая функция экранирования для предотвращения XSS
	return fmt.Sprintf("%q", param)
}

func main() {
	web.Router("/xss-01/BenchmarkTest00888", &BenchmarkTest00888Controller{})
	web.Run()
}
