package controllers

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02253Controller struct {
	web.Controller
}

func (c *BenchmarkTest02253Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02253Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02253")
	bar := doSomething(param)

	var cmd string
	var args []string
	osName := "Linux" // You might need a better way to detect the OS

	if strings.Contains(osName, "Windows") {
		args = []string{"cmd.exe", "/c", "echo ", bar}
	} else {
		cmd = "ls " + bar
		args = []string{"sh", "-c", cmd}
	}

	cmdExec := exec.Command(args[0], args[1:]...)
	output, err := cmdExec.CombinedOutput()
	if err != nil {
		fmt.Println("Problem executing cmdi - TestCase")
		c.Ctx.Output.Body([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}

	c.Ctx.Output.Body(output)
}

func doSomething(param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[1]

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
