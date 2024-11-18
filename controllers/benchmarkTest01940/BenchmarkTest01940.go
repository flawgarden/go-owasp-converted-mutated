package controllers

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"os/exec"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01940 struct {
	web.Controller
}

func (c *BenchmarkTest01940) Get() {
	c.Post()
}

func (c *BenchmarkTest01940) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Header("BenchmarkTest01940")
	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	cmd := ""
	osName := ""

	if strings.Contains(osName, "Windows") {
		cmd = "echo "
	}

	argsEnv := []string{"Foo=bar"}
	r := exec.Command(cmd+bar, argsEnv...)

	output, err := r.Output()
	if err != nil {
		fmt.Println("Problem executing cmdi - TestCase")
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	c.Ctx.Output.Body(output)
}

func doSomething(param string) string {
	var bar string
	if param != "" {
		bar = string(base64.StdEncoding.EncodeToString([]byte(param)))
	}
	return bar
}
