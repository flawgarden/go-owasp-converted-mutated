package controllers

import (
	"fmt"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01527 struct {
	web.Controller
}

func (c *BenchmarkTest01527) Get() {
	c.Post()
}

func (c *BenchmarkTest01527) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01527")
	if param == "" {
		param = ""
	}

	bar := testDoSomething(param)

	cmd := ""
	if isWindows() {
		cmd = getOSCommandString("echo")
	}

	out, err := exec.Command(cmd + bar).Output()
	if err != nil {
		fmt.Println("Problem executing cmdi - TestCase")
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	c.Data["json"] = string(out)
	c.ServeJSON()
}

func testDoSomething(param string) string {
	num := 106
	if (7*18)+num > 200 {
		return "This_should_always_happen"
	}
	return param
}

func isWindows() bool {
	return false // implement logic to check for Windows
}

func getOSCommandString(cmd string) string {
	return cmd // implement logic to return appropriate command string
}

func main() {
	web.Router("/cmdi-01/BenchmarkTest01527", &BenchmarkTest01527{})
	web.Run()
}
