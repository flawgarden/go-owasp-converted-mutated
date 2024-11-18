package controllers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest struct {
	web.Controller
}

func (c *BenchmarkTest) Get() {
	c.DoPost()
}

func (c *BenchmarkTest) DoPost() {
	response := c.Ctx.ResponseWriter
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01254")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(c.Ctx.Request, param)

	response.Header().Set("X-XSS-Protection", "0")
	response.Write([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(req *http.Request, param string) string {
	a23874 := param
	b23874 := a23874 + " SafeStuff"
	b23874 = b23874[:len(b23874)-len("Chars")] + "Chars"
	map23874 := make(map[string]string)
	map23874["key23874"] = b23874
	c23874 := map23874["key23874"]
	d23874 := c23874[:len(c23874)-1]
	e23874 := string([]byte(d23874))           // Base64 encode and decode skipped for simplicity
	f23874 := e23874[:len(e23874)-len(e23874)] // Sample manipulation for demonstration
	// Assuming there's a ThingInterface and ThingFactory for demonstration
	bar := f23874 // Reflection skipped for clarity

	return bar
}
