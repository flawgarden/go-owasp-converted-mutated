package controllers

import (
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type SqlInjectionVuln2Controller struct {
	web.Controller
}

func (c *SqlInjectionVuln2Controller) Get() {
	param := c.GetString("BenchmarkTest01259")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	output := fmt.Sprintf("Formatted like: %s.", bar)
	c.Ctx.ResponseWriter.Write([]byte(output))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := "safe!"
	mapping := make(map[string]interface{})
	mapping["keyA-26093"] = "a-Value"
	mapping["keyB-26093"] = param
	mapping["keyC"] = "another-Value"
	bar = mapping["keyB-26093"].(string)

	return bar
}
