package controllers

import (
	"fmt"
	"strconv"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01456 struct {
	web.Controller
}

func (c *BenchmarkTest01456) Get() {
	c.Post()
}

func (c *BenchmarkTest01456) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	names := c.Ctx.Request.URL.Query()
	for name, values := range names {
		for _, value := range values {
			if value == "BenchmarkTest01456" {
				param = name
				break
			}
		}
	}

	bar := doSomething(param)

	c.Ctx.Request.Context().Value("session").(map[string]interface{})["userid"] = bar

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Item: 'userid' with value: '%s' saved in session.", encodeForHTML(bar))))
}

func doSomething(param string) string {
	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}

func encodeForHTML(s string) string {
	return strconv.Quote(s) // Простой пример кодирования
}

func main() {
	web.Router("/trustbound-00/BenchmarkTest01456", &BenchmarkTest01456{})
	web.Run()
}
