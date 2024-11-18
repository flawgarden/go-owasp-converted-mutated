package controllers

import (
	"fmt"
	"os/exec"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

type BenchmarkTest02612 struct {
	beego.Controller
}

func (c *BenchmarkTest02612) Get() {
	c.Post()
}

func (c *BenchmarkTest02612) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Input.URI()
	paramval := "BenchmarkTest02612="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramval) - 1
	}
	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02612")))
		return
	}
	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := len(queryString)

	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}
	param = decode(param)

	bar := doSomething(param)

	cmd := "your-command-here" // replace with your command
	args := []string{cmd}
	argsEnv := []string{bar}

	r := exec.Command(args[0], args[1:]...)
	r.Env = append(r.Env, argsEnv...)

	output, err := r.Output()
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Problem executing cmdi - TestCase: %s", err.Error())))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	// Consider implementing your own logic here.
	return param
}

func decode(s string) string {
	// Implement your decoding logic if necessary.
	return s
}
