package controllers

import (
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01618Controller struct {
	web.Controller
}

func (c *BenchmarkTest01618Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01618Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01618")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(c.Ctx.Request, param)

	c.Ctx.Request.Header.Set("userid", bar)
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Item: 'userid' with value: '%s' saved in session.", bar)))
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	// Здесь вы можете воспользоваться интерфейсом или другой логикой
	// но оставим просто возврат param для примера
	return param
}
