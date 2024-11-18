package controllers

import (
	"encoding/base64"
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

type BenchmarkTest01927 struct {
	beego.Controller
}

func (c *BenchmarkTest01927) Get() {
	c.Post()
}

func (c *BenchmarkTest01927) Post() {
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	param := c.Ctx.Request.Header.Get("Referer")
	if param != "" {
		param = decodeURL(param)
	}
	bar := doSomething(c.Ctx.Request, param)
	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Ctx.WriteString(bar)
}

func doSomething(request *http.Request, param string) string {
	bar := ""
	if param != "" {
		decodedParam, _ := base64.StdEncoding.DecodeString(param)
		bar = string(decodedParam)
	}
	return bar
}

func decodeURL(value string) string {
	decodedValue, _ := url.QueryUnescape(value)
	return decodedValue
}
