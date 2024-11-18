package controllers

import (
	"net/http"
	"net/url"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01192Controller struct {
	web.Controller
}

func (c *BenchmarkTest01192Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01192Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["BenchmarkTest01192"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := new(test).doSomething(c.Ctx.Request, param)

	cmd := "your-command-here" // Replace with your command
	args := []string{cmd}
	argsEnv := []string{bar}

	r := exec.Command(args[0], args[1:]...)
	r.Env = append(r.Env, argsEnv...)

	output, err := r.CombinedOutput()
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Problem executing cmdi - TestCase: " + err.Error()))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

type test struct{}

func (t *test) doSomething(request *http.Request, param string) string {
	bar := ""
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}
