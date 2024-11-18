package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

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

type BenchmarkTest01175Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01175Controller) Get() {
	c.postRequest()
}

func (c *BenchmarkTest01175Controller) Post() {
	c.postRequest()
}

func (c *BenchmarkTest01175Controller) postRequest() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if referer := c.Ctx.Request.Header.Get("Referer"); referer != "" {
		param = referer
	}

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(c.Ctx.Request, param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	output := fmt.Sprintf("Formatted like: %s and %s.", "a", bar)
	c.Ctx.ResponseWriter.Write([]byte(output))
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	bar := strings.ReplaceAll(param, "&", "&amp;")
	bar = strings.ReplaceAll(bar, "<", "&lt;")
	bar = strings.ReplaceAll(bar, ">", "&gt;")
	return bar
}
