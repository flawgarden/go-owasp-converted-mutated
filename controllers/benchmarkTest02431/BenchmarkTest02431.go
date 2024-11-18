package controllers

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02431Controller struct {
	web.Controller
}

func (c *BenchmarkTest02431Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02431Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02431")
	if param == "" {
		param = ""
	}

	bar := doSomething(c.Ctx.Request, param)

	cmd := getInsecureOSCommandString()
	args := []string{cmd}
	argsEnv := []string{bar}

	r := exec.Command(args[0], args[1:]...)
	r.Env = argsEnv

	output, err := r.CombinedOutput()
	if err != nil {
		fmt.Println("Problem executing cmdi - TestCase")
		c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(req *http.Request, param string) string {
	bar := "safe!"
	map58889 := make(map[string]interface{})
	map58889["keyA-58889"] = "a-Value"
	map58889["keyB-58889"] = param
	map58889["keyC"] = "another-Value"
	bar = fmt.Sprintf("%v", map58889["keyB-58889"])
	return bar
}

func getInsecureOSCommandString() string {
	// Replace with actual implementation to retrieve command string
	return "echo"
}
