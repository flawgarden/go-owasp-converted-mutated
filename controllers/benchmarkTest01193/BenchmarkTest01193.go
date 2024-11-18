package controllers

import (
	"net/http"
	"net/url"

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

type BenchmarkTest01193Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01193Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01193Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if headers := c.Ctx.Request.Header["BenchmarkTest01193"]; len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(param)

	var cmd string
	var args []string
	if isWindows() {
		args = []string{"cmd.exe", "/c", "echo " + bar}
	} else {
		cmd = bar
		args = []string{"sh", "-c", cmd}
	}

	output, err := runCommand(args)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(output))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	var bar string
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func isWindows() bool {
	return false // replace with actual check for OS type
}

func runCommand(args []string) (string, error) {
	// Implement the command execution logic here
	return "", nil
}
