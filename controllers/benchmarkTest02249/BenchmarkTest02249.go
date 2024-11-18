package controllers

import (
	"fmt"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02249Controller struct {
	web.Controller
}

func (c *BenchmarkTest02249Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest02249Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest02249Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02249")
	bar := doSomething(param)

	var cmd string
	if isWindows() {
		cmd = "echo " + bar
	}

	out, err := exec.Command("cmd", "/C", cmd).CombinedOutput()
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Problem executing cmdi - TestCase: %s", err.Error())))
		return
	}

	c.Ctx.ResponseWriter.Write(out)
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = append(valuesList[:0], valuesList[1:]...) // remove the 1st safe value
		bar = valuesList[0]                                    // get the param value
	}
	return bar
}

func isWindows() bool {
	return false // replace with actual check for Windows if needed
}
