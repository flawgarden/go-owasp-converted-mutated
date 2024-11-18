package controllers

import (
	"database/sql"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02433 struct {
	web.Controller
}

func init() {
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func (c *BenchmarkTest02433) Get() {
	c.Post()
}

func (c *BenchmarkTest02433) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02433")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	var cmd string
	var a1, a2 string
	var args []string
	if isWindows() {
		a1 = "cmd.exe"
		a2 = "/c"
		cmd = "echo "
		args = []string{a1, a2, cmd + bar}
	} else {
		a1 = "sh"
		a2 = "-c"
		cmd = "ls "
		args = []string{a1, a2, cmd + bar}
	}

	if err := executeCommand(args); err != nil {
		c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}
}

func doSomething(param string) string {
	// Simulated functionality for processing
	return param
}

func isWindows() bool {
	return false // Replace with actual check
}

func executeCommand(args []string) error {
	r := exec.Command(args[0], args[1:]...)
	_, err := r.Output()
	return err
}
