package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00907Controller struct {
	web.Controller
}

func (c *BenchmarkTest00907Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00907Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00907")

	var bar string
	num := 106

	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	var cmd string
	var a1 string
	var a2 string
	var args []string
	osName := "linux" // Replace with actual OS detection if needed

	if osName == "windows" {
		a1 = "cmd.exe"
		a2 = "/c"
		cmd = "echo"
		args = []string{a1, a2, cmd, bar}
	} else {
		a1 = "sh"
		a2 = "-c"
		cmd = "ping -c1 "
		args = []string{a1, a2, cmd + bar}
	}

	if err := exec.Command(args[0], args[1:]...).Run(); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.Output.Body([]byte(fmt.Sprintf("Problem executing cmd - TestCase: %s", err.Error())))
		return
	}

	output, _ := json.Marshal("Command executed successfully")
	c.Ctx.Output.Body(output)
}
