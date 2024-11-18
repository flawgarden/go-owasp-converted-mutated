package controllers

import (
	"os/exec"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01780Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01780Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01780Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest01780")
	bar := new(Test).doSomething(param)

	var a1, a2 string
	if isWindows() {
		a1 = "cmd.exe"
		a2 = "/c"
	} else {
		a1 = "sh"
		a2 = "-c"
	}

	args := []string{a1, a2, "echo " + bar}
	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.WriteString(string(output))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := "safe!"
	map90655 := make(map[string]interface{})
	map90655["keyA-90655"] = "a_Value"
	map90655["keyB-90655"] = param
	map90655["keyC"] = "another_Value"

	bar = map90655["keyB-90655"].(string)
	bar = map90655["keyA-90655"].(string)

	return bar
}

func isWindows() bool {
	return exec.Command("cmd", "/c", "ver").Run() == nil
}
