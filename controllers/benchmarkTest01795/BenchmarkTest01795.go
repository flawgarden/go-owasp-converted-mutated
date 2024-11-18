package controllers

import (
	"fmt"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01795Controller struct {
	web.Controller
}

func (c *BenchmarkTest01795Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01795Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01795")

	bar := new(Test).doSomething(param)

	cmd := getInsecureOSCommandString()
	args := []string{cmd}
	argsEnv := []string{bar}

	r := exec.Command(args[0], args[1:]...)
	r.Env = append(r.Env, argsEnv...)

	output, err := r.Output()
	if err != nil {
		fmt.Println("Problem executing cmdi - TestCase")
		c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Ctx.ResponseWriter.Write(output)
}

func getInsecureOSCommandString() string {
	return "echo" // пример команды
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value

		bar = valuesList[0] // get the param value
	}
	return bar
}
