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

type BenchmarkTest01550 struct {
	beego.Controller
}

func (c *BenchmarkTest01550) Get() {
	c.Post()
}

func (c *BenchmarkTest01550) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest01550")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	c.Ctx.Request.Context().Value("userid").(http.ResponseWriter).Write([]byte(fmt.Sprintf("Item: 'userid' with value: '%s' saved in session.", bar)))
}

func doSomething(param string) string {
	a := param
	b := fmt.Sprintf("%s SafeStuff", a)
	b = b[:len(b)-1] + "Chars"
	c := make(map[string]interface{})
	c["key"] = b
	d := c["key"].(string)
	e := d[:len(d)-1]
	f := e // Base64 encoding/decoding is omitted for simplicity
	return f
}
