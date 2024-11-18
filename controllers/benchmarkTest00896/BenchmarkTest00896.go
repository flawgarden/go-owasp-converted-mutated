package controllers

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00896Controller struct {
	web.Controller
}

func (c *BenchmarkTest00896Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest00896Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest00896Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00896")
	bar := ""

	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	var argList []string
	osName := "" // Replace with actual OS detection if needed
	if strings.Contains(osName, "Windows") {
		argList = append(argList, "cmd.exe", "/c")
	} else {
		argList = append(argList, "sh", "-c")
	}
	argList = append(argList, "echo "+bar)

	cmd := exec.Command(argList[0], argList[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Problem executing command:", err)
		c.Ctx.WriteString("Error executing command")
		return
	}

	c.Ctx.WriteString(string(output))
}
