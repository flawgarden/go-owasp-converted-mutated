package controllers

import (
	"net/http"
	"net/url"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02069Controller struct {
	web.Controller
}

func (c *BenchmarkTest02069Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02069Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["BenchmarkTest02069"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(c.Ctx.Request, param)

	cmd := ""
	if isWindows() {
		cmd = "cmd /C echo "
	}

	argsEnv := []string{"Foo=bar"}
	r := exec.Command(cmd+bar, argsEnv...)
	output, err := r.Output()

	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Problem executing cmdi - TestCase: " + err.Error()))
		return
	}

	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(req *http.Request, param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	return bar
}

func isWindows() bool {
	return false // Замените на логику проверки ОС, если нужно
}
