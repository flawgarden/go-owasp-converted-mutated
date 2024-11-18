package controllers

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01269Controller struct {
	web.Controller
}

func (c *BenchmarkTest01269Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01269Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01269")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(c.Ctx.Request, param)

	var a1, a2 string
	if isWindows() {
		a1 = "cmd.exe"
		a2 = "/c"
	} else {
		a1 = "sh"
		a2 = "-c"
	}
	args := []string{a1, a2, "echo " + bar}

	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Problem executing command")
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}
	c.Ctx.Output.Body(output)
}

func isWindows() bool {
	return os.Getenv("OS") == "Windows_NT"
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	bar := param
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	}
	return bar
}
