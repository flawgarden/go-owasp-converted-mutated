package controllers

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02149Controller struct {
	web.Controller
}

func (c *BenchmarkTest02149Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02149Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02149")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	var cmd string
	var a1, a2 string
	var args []string
	osName := detectOS()

	if strings.Contains(osName, "Windows") {
		a1 = "cmd.exe"
		a2 = "/c"
		cmd = "echo "
		args = []string{a1, a2, cmd + bar}
	} else {
		a1 = "sh"
		a2 = "-c"
		cmd = "ls "
		args = []string{a1, a2, cmd + bar}
	}

	argsEnv := []string{"foo=bar"}

	r := exec.Command(args[0], args[1:]...)
	r.Env = append(r.Env, argsEnv...)
	output, err := r.CombinedOutput()
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	bar := ""

	num := 106

	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}

func detectOS() string {
	return "Linux" // или "Windows" в зависимости от среды
}
