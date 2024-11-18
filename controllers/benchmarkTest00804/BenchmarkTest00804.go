package controllers

import (
	"fmt"
	"strings"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00804 struct {
	web.Controller
}

func (c *BenchmarkTest00804) Get() {
	c.Post()
}

func (c *BenchmarkTest00804) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Request.URL.RawQuery
	paramVal := "BenchmarkTest00804="
	paramLoc := strings.Index(queryString, paramVal)
	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest00804")))
		return
	}

	param := queryString[paramLoc+len(paramVal):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramVal) : ampersandLoc+paramLoc]
	}

	param = param // Decode if necessary

	var sbxyz12823 strings.Builder
	sbxyz12823.WriteString(param)
	bar := sbxyz12823.String() + "_SafeStuff"

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}
