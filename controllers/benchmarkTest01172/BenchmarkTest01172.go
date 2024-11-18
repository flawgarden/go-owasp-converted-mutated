package controllers

import (
	"fmt"
	"net/http"
	"net/url"

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

type BenchmarkTest01172Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01172Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01172Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	referer := c.Ctx.Request.Header.Get("Referer")
	if referer != "" {
		param = referer
	}

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(c.Ctx.Request, param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	fmt.Fprintf(c.Ctx.ResponseWriter, "Formatted like: %s and %s.", "a", bar)
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	bar := "safe!"
	map35084 := make(map[string]interface{})
	map35084["keyA-35084"] = "a-Value"
	map35084["keyB-35084"] = param
	map35084["keyC"] = "another-Value"
	bar = map35084["keyB-35084"].(string)

	return bar
}
