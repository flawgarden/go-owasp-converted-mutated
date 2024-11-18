package controllers

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01529Controller struct {
	web.Controller
}

func (c *BenchmarkTest01529Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01529Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest01529")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(c.Ctx.Request, param)

	var args []string
	osName := "" // replace with actual os.Name retrieval logic

	if strings.Contains(osName, "Windows") {
		args = append(args, "cmd.exe", "/c", "echo "+bar)
	} else {
		args = append(args, "sh", "-c", "ping -c1 "+bar)
	}

	r := exec.Command(args[0], args[1:]...)

	output, err := r.CombinedOutput()
	if err != nil {
		c.Ctx.Output.Body([]byte(fmt.Sprintf("Problem executing cmdi - TestCase: %s", err.Error())))
		return
	}

	c.Ctx.Output.Body(output)
}

type Test struct{}

func (t *Test) doSomething(req *http.Request, param string) string {
	bar := param
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	}
	return bar
}
