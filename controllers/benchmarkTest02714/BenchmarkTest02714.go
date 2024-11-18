package controllers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02714Controller struct {
	web.Controller
}

func (c *BenchmarkTest02714Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest02714Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest02714Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Query("BenchmarkTest02714")
	bar := doSomething(c.Ctx.Request, param)

	cmd := getInsecureOSCommandString()
	argsEnv := []string{bar}

	if err := executeCommand(cmd, argsEnv); err != nil {
		c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}
}

func doSomething(req *http.Request, param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}

func getInsecureOSCommandString() string {
	// Add logic to fetch the insecure OS command string
	return "insecure_command"
}

func executeCommand(cmd string, args []string) error {
	// Implement the command execution logic here
	return nil
}
