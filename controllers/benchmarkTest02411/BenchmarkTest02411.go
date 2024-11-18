package controllers

import (
	"net/http"
	"os/exec"
	"runtime"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02411 struct {
	web.Controller
}

func (c *BenchmarkTest02411) Get() {
	c.Post()
}

func (c *BenchmarkTest02411) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest02411")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	var cmd string
	if isWindows() {
		cmd = "cmd.exe /c echo " + bar
	} else {
		cmd = "sh -c echo " + bar
	}

	if err := executeCommand(cmd, c.Ctx.ResponseWriter); err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}

func isWindows() bool {
	return runtime.GOOS == "windows"
}

func executeCommand(cmd string, response http.ResponseWriter) error {
	command := exec.Command("sh", "-c", cmd)
	output, err := command.CombinedOutput()
	if err != nil {
		return err
	}
	response.Write(output)
	return nil
}
