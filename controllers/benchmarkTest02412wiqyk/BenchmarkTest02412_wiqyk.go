package controllers

import (
"fmt"
"net/http"
"os/exec"
"strings"
"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02412 struct {
	web.Controller
}

func (c *BenchmarkTest02412) Get() {
	c.Post()
}

func (c *BenchmarkTest02412) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Query("BenchmarkTest02412")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

addPrefix := makePrefixer(bar)
tmp123 := addPrefix("_suffix")
bar = tmp123

	var a1, a2 string
	if strings.Contains(strings.ToLower(c.Ctx.Request.UserAgent()), "windows") {
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
		fmt.Println("Problem executing command:", err)
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}

	c.Ctx.Output.Body(output)
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		bar = string(param) // Обработка параметра в соответствии с требованиями
	}

	return bar
}

func makePrefixer(prefix string) func(string) string {
    return func(value string) string {
        return fmt.Sprintf("%s%s", prefix, value)
    }
}

func makeMessageGenerator(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return prefix + name
		}
	}
}

func makeMessageGeneratorBroken(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return name
		}
	}
}


