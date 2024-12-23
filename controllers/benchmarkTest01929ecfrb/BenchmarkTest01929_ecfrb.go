package controllers

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01929Controller struct {
	web.Controller
}

func (c *BenchmarkTest01929Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01929Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Request.Header.Get("BenchmarkTest01929")

var a12341 BaseBinaryOpClass = &DerivedBinaryOpClassDefault{}
param = a12341.VirtualCall(param, param)

	bar := doSomething(param)

	var a1, a2 string
	if isWindowsOS() {
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
		fmt.Println("Problem executing command")
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	bar := "safe!"
	map44 := make(map[string]interface{})
	map44["keyA-44"] = "a-Value"
	map44["keyB-44"] = param
	map44["keyC"] = "another-Value"
	bar = fmt.Sprintf("%v", map44["keyB-44"])
	return bar
}

func isWindowsOS() bool {
	return exec.Command("cmd", "/C", "ver").Run() == nil
}