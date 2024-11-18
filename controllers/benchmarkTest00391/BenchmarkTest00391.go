package controllers

import (
	"database/sql"
	"html"

	"github.com/beego/beego/v2/server/web"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00391 struct {
	web.Controller
}

func (c *BenchmarkTest00391) Get() {
	c.Post()
}

func (c *BenchmarkTest00391) Post() {
	response := c.Ctx.ResponseWriter
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")
	response.Header().Set("X-XSS-Protection", "0")

	param := c.GetString("BenchmarkTest00391")
	if param == "" {
		param = ""
	}
	bar := html.EscapeString(param)

	response.Write([]byte(bar))
}

func init() {
	_, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
}
