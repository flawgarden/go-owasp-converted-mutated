package controllers

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02252Controller struct {
	web.Controller
}

func (c *BenchmarkTest02252Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02252Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02252")
	bar := doSomething(param)

	var cmd string
	var args []string
	osName := os.Getenv("OS")

	if osName != "" && osName == "Windows_NT" {
		args = []string{"cmd.exe", "/c", "echo " + bar}
	} else {
		args = []string{"sh", "-c", "ls " + bar}
	}

	cmd = args[0]
	args = args[1:]

	cmdProcess := exec.Command(cmd, args...)
	output, err := cmdProcess.CombinedOutput()

	if err != nil {
		fmt.Println("Problem executing cmdi - TestCase")
		c.Ctx.ResponseWriter.Write([]byte(escapeHTML(err.Error())))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	bar := param
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	}
	return bar
}

func escapeHTML(input string) string {
	return input // Implement HTML escape as needed
}
