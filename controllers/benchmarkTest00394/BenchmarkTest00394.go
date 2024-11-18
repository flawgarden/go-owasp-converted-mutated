package controllers

import (
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00394Controller struct {
	web.Controller
}

func (c *BenchmarkTest00394Controller) Get() {
	c.post()
}

func (c *BenchmarkTest00394Controller) Post() {
	c.post()
}

func (c *BenchmarkTest00394Controller) post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00394")
	if param == "" {
		param = ""
	}

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	if bar != "" {
		c.Ctx.ResponseWriter.Write([]byte(bar))
	}
}
