package controllers

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02156Controller struct {
	web.Controller
}

func (c *BenchmarkTest02156Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02156Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02156")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	cmd := getInsecureOSCommandString()
	argsEnv := []string{bar}
	r := exec.Command(cmd, argsEnv...)
	output, err := r.CombinedOutput()
	if err != nil {
		fmt.Println("Problem executing cmdi - TestCase")
		c.Ctx.ResponseWriter.Write([]byte(htmlEscape(err.Error())))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	var bar string
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func getInsecureOSCommandString() string {
	return "your-command-here" // Замените на вашу команду
}

func htmlEscape(input string) string {
	return strings.ReplaceAll(strings.ReplaceAll(input, "&", "&amp;"), "<", "&lt;")
}
