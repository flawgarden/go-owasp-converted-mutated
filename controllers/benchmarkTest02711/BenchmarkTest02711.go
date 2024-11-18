package controllers

import (
	"fmt"
	"net/http"

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

type BenchmarkTest02711Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02711Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02711Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest02711")

	bar := doSomething(param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   true,
		HttpOnly: true,
		Path:     c.Ctx.Request.RequestURI,
	}
	http.SetCookie(c.Ctx.ResponseWriter, &cookie)

	output := fmt.Sprintf("Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: true", bar)
	c.Ctx.ResponseWriter.Write([]byte(output))
}

func doSomething(param string) string {
	return param
}
