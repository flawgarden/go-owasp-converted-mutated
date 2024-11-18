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

type BenchmarkTest01341 struct {
	beego.Controller
}

func (c *BenchmarkTest01341) Get() {
	c.handleRequest()
}

func (c *BenchmarkTest01341) Post() {
	c.handleRequest()
}

func (c *BenchmarkTest01341) handleRequest() {
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	param := c.GetString("BenchmarkTest01341")

	bar := new(Test).doSomething(param)

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Formatted like: %s and %s.", "a", bar)))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	a := param
	b := a + " SafeStuff"
	return b[:len(b)-1] // extracting part of the string
}
