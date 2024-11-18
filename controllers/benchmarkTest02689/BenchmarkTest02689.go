package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02689Controller struct {
	web.Controller
}

func (c *BenchmarkTest02689Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest02689Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest02689Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	id := c.GetString("BenchmarkTest02689")

	bar := doSomething(id)

	output, err := json.Marshal(bar)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}

func main() {
	web.Router("/xss-05/BenchmarkTest02689", &BenchmarkTest02689Controller{})
	web.Run()
}
