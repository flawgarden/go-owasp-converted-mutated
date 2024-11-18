package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01493 struct {
	web.Controller
}

func (c *BenchmarkTest01493) Get() {
	c.Post()
}

func (c *BenchmarkTest01493) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01493")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(c.Ctx.Request, param)

	fileTarget := fmt.Sprintf("%s/%s", os.Getenv("TESTFILES_DIR"), bar)
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", fileTarget)))
	if _, err := os.Stat(fileTarget); err == nil {
		c.Ctx.ResponseWriter.Write([]byte(" And file already exists."))
	} else {
		c.Ctx.ResponseWriter.Write([]byte(" But file doesn't exist yet."))
	}
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map71009 := make(map[string]interface{})
	map71009["keyA-71009"] = "a_Value"
	map71009["keyB-71009"] = param
	map71009["keyC"] = "another_Value"
	bar = map71009["keyB-71009"].(string)
	bar = map71009["keyA-71009"].(string)

	return bar
}

func main() {
	web.Router("/pathtraver-01/BenchmarkTest01493", &BenchmarkTest01493{})
	web.Run()
}
