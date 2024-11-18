package controllers

import (
	"os/exec"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00743Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00743Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest00743Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest00743Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00743", "")
	var bar string
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

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

	cmd := "your_command_here" // Specify your command
	argsEnv := strings.Split(bar, " ")
	r := exec.Command(cmd, argsEnv...)

	output, err := r.CombinedOutput()
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Problem executing cmdi - TestCase: " + err.Error()))
		return
	}
	c.Ctx.WriteString(string(output))
}
