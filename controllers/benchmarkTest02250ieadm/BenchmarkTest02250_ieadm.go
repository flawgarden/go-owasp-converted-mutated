package controllers

import (
"fmt"
"os/exec"
"github.com/beego/beego/v2/server/web"
"sync"
)

type BenchmarkTest02250 struct {
	web.Controller
}

func (c *BenchmarkTest02250) Get() {
	c.Post()
}

func (c *BenchmarkTest02250) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02250", "")

message123 := make(chan string, 1)
message123 <- param

var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    _ = <-message123
    message123 <- "constant_string"
}()

wg.Wait()

param = <-message123

	bar := doSomething(param)

	cmd := ""
	if isWindows() {
		cmd = getOSCommandString("echo ")
	}

	out, err := exec.Command("sh", "-c", cmd+bar).Output()
	if err != nil {
		fmt.Println("Problem executing cmdi - TestCase")
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	c.Ctx.Output.Body(out)
}

func doSomething(param string) string {
	bar := "safe!"
	map94176 := make(map[string]interface{})
	map94176["keyA-94176"] = "a-Value"
	map94176["keyB-94176"] = param
	map94176["keyC"] = "another-Value"
	bar = map94176["keyB-94176"].(string)
	return bar
}

func isWindows() bool {
	return exec.Command("cmd").Run() == nil
}

func getOSCommandString(cmd string) string {
	return cmd
}
