package controllers

import (
	"os/exec"
	"runtime"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02698Controller struct {
	web.Controller
}

func init() {
	web.Router("/cmdi-03/BenchmarkTest02698", &BenchmarkTest02698Controller{})
}

func (c *BenchmarkTest02698Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest02698Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02698")
	bar := doSomething(param)

	var argList []string
	if isWindows() {
		argList = []string{"cmd.exe", "/c", "echo " + bar}
	} else {
		argList = []string{"sh", "-c", "echo " + bar}
	}

	_, err := exec.Command(argList[0], argList[1:]...).Output()
	if err != nil {
		panic(err)
	}
}

func doSomething(param string) string {
	// Assume this function interacts with an external system and returns a processed string
	return param // Simplified for demonstration
}

func isWindows() bool {
	return runtime.GOOS == "windows"
}
