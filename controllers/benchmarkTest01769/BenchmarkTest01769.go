package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01769Controller struct {
	web.Controller
}

func (c *BenchmarkTest01769Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01769Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	param := c.GetString("BenchmarkTest01769")
	bar := doSomething(param)

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(bar, "a", "b")))
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}
	return bar
}
