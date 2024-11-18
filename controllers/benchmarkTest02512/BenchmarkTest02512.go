package controllers

import (
	"fmt"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02512Controller struct {
	web.Controller
}

func (c *BenchmarkTest02512Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02512Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02512")
	bar := doSomething(param)

	cmd := "your_command_here" // Replace with the actual command you want to execute
	args := []string{cmd}
	argsEnv := []string{bar}

	r := exec.Command(args[0], args[1:]...)
	r.Env = append(r.Env, argsEnv...)

	out, err := r.Output()
	if err != nil {
		c.Ctx.Output.Body([]byte(fmt.Sprintf("Problem executing cmdi - TestCase: %s", err.Error())))
		return
	}

	c.Ctx.Output.Body(out)
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}
