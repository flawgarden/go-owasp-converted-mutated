package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01337Controller struct {
	web.Controller
}

func (c *BenchmarkTest01337Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest01337Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest01337Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01337")

	bar := (&Test{}).doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	_, err := fmt.Fprintf(c.Ctx.ResponseWriter, bar, "a", "b")
	if err != nil {
		panic(err)
	}
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	sbxyz24804 := param + "_SafeStuff"
	return sbxyz24804
}
