package controllers

import (
	"os/exec"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02344Controller struct {
	web.Controller
}

func (c *BenchmarkTest02344Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02344Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := ""
	flag := true

	names := c.Ctx.Request.URL.Query()
	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest02344" {
					param = name
					flag = false
				}
			}
		}
	}

	bar := doSomething(param)

	cmd := getInsecureOSCommandString()
	argsEnv := []string{bar}
	cmdOut, err := exec.Command(cmd, argsEnv...).CombinedOutput()
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Problem executing cmdi - TestCase"))
		c.Ctx.ResponseWriter.Write([]byte(encodeForHTML(err.Error())))
		return
	}
	c.Ctx.ResponseWriter.Write(cmdOut)
}

func doSomething(param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
}

func getInsecureOSCommandString() string {
	return "your_command_here" // Replace with the actual command
}

func encodeForHTML(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}
