package controllers

import (
	"encoding/json"
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

type BenchmarkTest02709Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02709Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02709Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02709")
	bar := doSomething(param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   false,
		HttpOnly: true,
		Path:     c.Ctx.Input.URI(),
	}
	http.SetCookie(c.Ctx.ResponseWriter, &cookie)

	output, err := json.Marshal(fmt.Sprintf("Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: false", bar))
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	bar := param
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	}
	return bar
}
