package controllers

import (
	"fmt"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01794Controller struct {
	web.Controller
}

func (c *BenchmarkTest01794Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01794Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest01794")

	bar := doSomething(param)

	var args []string
	if isWindows() {
		args = []string{"cmd.exe", "/c", "echo", bar}
	} else {
		args = []string{"sh", "-c", "ls " + bar}
	}

	if err := executeCommand(args); err != nil {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}
}

func doSomething(param string) string {
	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}

func isWindows() bool {
	return false // Simple check for Windows or Unix
}

func executeCommand(args []string) error {
	cmd := exec.Command(args[0], args[1:]...)
	return cmd.Run()
}
