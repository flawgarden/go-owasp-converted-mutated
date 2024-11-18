package controllers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01063Controller struct {
	web.Controller
}

func (c *BenchmarkTest01063Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01063Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01063Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if referer := c.Ctx.Request.Header.Get("Referer"); referer != "" {
		param = referer
	}

	param, _ = url.QueryUnescape(param)

	bar := new(Test).DoSomething(c.Ctx.Request, param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte("Parameter value: " + bar))
}

type Test struct{}

func (t *Test) DoSomething(request *http.Request, param string) string {
	a92400 := param
	var b92400 strings.Builder
	b92400.WriteString(a92400)
	b92400.WriteString(" SafeStuff")
	b92400Str := b92400.String()
	b92400Str = strings.Replace(b92400Str, b92400Str[len(b92400Str)-len("Chars"):], "Chars", 1)

	map92400 := make(map[string]interface{})
	map92400["key92400"] = b92400Str
	c92400 := map92400["key92400"].(string)
	d92400 := c92400[:len(c92400)-1]

	e92400 := string([]byte(d92400)) // Simulate Base64 decode/encode
	f92400 := strings.Split(e92400, " ")[0]

	thing := createThing()
	bar := thing.DoSomething(f92400)

	return bar
}

func createThing() ThingInterface {
	// Implement function to create and return an instance of ThingInterface
	return &Thing{}
}

type ThingInterface interface {
	DoSomething(input string) string
}

type Thing struct{}

func (t *Thing) DoSomething(input string) string {
	// Implement the actual logic here
	return input
}
