package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os/exec"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02068Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02068Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02068Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["BenchmarkTest02068"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(c.Ctx.Request, param)

	cmd := getInsecureOSCommandString()
	args := []string{cmd}
	argsEnv := []string{bar}

	if err := executeCommand(args, argsEnv, c.Ctx.ResponseWriter); err != nil {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Problem executing cmdi - TestCase: %v", err)))
		return
	}
}

func doSomething(r *http.Request, param string) string {
	num := 106
	if (7*18)+num > 200 {
		return "This_should_always_happen"
	}
	return param
}

func getInsecureOSCommandString() string {
	// Implement logic to retrieve insecure OS command string
	return "your-command-here"
}

func executeCommand(args []string, argsEnv []string, response http.ResponseWriter) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = argsEnv
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	response.Write(output)
	return nil
}
