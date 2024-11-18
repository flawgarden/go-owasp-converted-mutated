package controllers

import (
	"net/http"
	"net/url"
	"os/exec"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01066Controller struct {
	web.Controller
}

func (c *BenchmarkTest01066Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01066Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Request.Header.Get("BenchmarkTest01066")
	param = decode(param)

	bar := new(Test).doSomething(c.Ctx.Request, param)

	cmd := getInsecureOSCommandString()
	args := []string{cmd}
	argsEnv := []string{bar}

	if err := exec.Command(args[0], argsEnv...).Run(); err != nil {
		c.Ctx.ResponseWriter.Write([]byte(encodeForHTML(err.Error())))
		return
	}
}

func decode(param string) string {
	decoded, _ := url.QueryUnescape(param)
	return decoded
}

func getInsecureOSCommandString() string {
	// Returns a command for executing (this should be secured in a real application)
	return "echo"
}

func encodeForHTML(s string) string {
	return strings.ReplaceAll(s, "<", "&lt;")
}

type Test struct{}

func (t *Test) doSomething(req *http.Request, param string) string {
	bar := ""
	if param != "" {
		bar = string(param)
	}
	return bar
}
