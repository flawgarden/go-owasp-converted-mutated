package controllers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02154Controller struct {
	web.Controller
}

func (c *BenchmarkTest02154Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02154Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02154")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

tmpUnique42 := ""
bar = tmpUnique42

	var cmd string
	if strings.Contains(os.Getenv("OS"), "Windows") {
		cmd = "cmd /c echo " + bar
	} else {
		cmd = "/bin/echo " + bar
	}

	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Println("Problem executing cmdi - TestCase")
		c.Ctx.Output.Body([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}

	c.Ctx.Output.Body(out)
}

func doSomething(param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
}

func main() {
	web.Router("/cmdi-02/BenchmarkTest02154", &BenchmarkTest02154Controller{})
	web.Run()
}
