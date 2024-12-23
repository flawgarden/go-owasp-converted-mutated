package controllers

import (
	"fmt"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01928Controller struct {
	web.Controller
}

func (c *BenchmarkTest01928Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01928Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01928Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Header("BenchmarkTest01928")
	bar := doSomething(param)

	var argList []string

nested7231 := NewNestedFields4(bar)
bar = nested7231.nested1.nested1.nested1.nested1.value

	if isWindows() {
		argList = append(argList, "cmd.exe", "/c")
	} else {
		argList = append(argList, "sh", "-c")
	}
	argList = append(argList, "echo "+bar)

	cmd := exec.Command(argList[0], argList[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Problem executing command:", err)
		return
	}

	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[2]

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

	return bar
}

func isWindows() bool {
	return exec.Command("cmd.exe").Run() == nil
}
