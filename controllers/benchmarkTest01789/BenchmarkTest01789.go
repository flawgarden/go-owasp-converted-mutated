package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01789Controller struct {
	web.Controller
}

func (c *BenchmarkTest01789Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01789Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.Ctx.Input.Query("BenchmarkTest01789")
	bar := testDoSomething(param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   false,
		HttpOnly: true,
		Path:     c.Ctx.Request.URL.Path,
	}
	http.SetCookie(c.Ctx.ResponseWriter, &cookie)

	output, err := json.Marshal(fmt.Sprintf("Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: false", bar))
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func testDoSomething(param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
}
