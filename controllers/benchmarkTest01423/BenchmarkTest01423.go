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

type BenchmarkTest01423 struct {
	beego.Controller
}

func (c *BenchmarkTest01423) Get() {
	c.Post()
}

func (c *BenchmarkTest01423) Post() {
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	param := extractParam(c.Ctx.Request)
	bar := new(Test).doSomething(c.Ctx.Request, param)
	obj := []interface{}{bar, "b"}
	fmt.Fprintf(c.Ctx.ResponseWriter, "Formatted like: %s and %s.", obj...)
}

func extractParam(request *http.Request) string {
	names := request.URL.Query()
	param := ""
	for name := range names {
		values := names[name]
		for _, value := range values {
			if value == "BenchmarkTest01423" {
				param = name
				break
			}
		}
		if param != "" {
			break
		}
	}
	return param
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	return param
}
