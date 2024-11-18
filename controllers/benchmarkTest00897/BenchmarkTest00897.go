package controllers

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00897Controller struct {
	web.Controller
}

func (c *BenchmarkTest00897Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00897Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00897")
	guess := "ABC"
	switchTarget := guess[2]

	var bar string
	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	var a1, a2 string
	if os := os.Getenv("OS"); os != "" && os == "Windows_NT" {
		a1 = "cmd.exe"
		a2 = "/c"
	} else {
		a1 = "sh"
		a2 = "-c"
	}
	args := []string{a1, a2, "echo " + bar}

	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Problem executing command Test Case")
		c.Ctx.WriteString(err.Error())
		return
	}
	c.Ctx.WriteString(string(output))
}
