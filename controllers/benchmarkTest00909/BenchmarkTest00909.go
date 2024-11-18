package controllers

import (
	"database/sql"
	"fmt"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	// Initialize database connection
	sql.Open("mysql", source)
}

type BenchmarkTest00909Controller struct {
	web.Controller
}

func (c *BenchmarkTest00909Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00909Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00909")
	var bar string

	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	cmd := "your-command-here"
	args := []string{cmd}
	argsEnv := []string{bar}

	r := exec.Command(args[0], args[1:]...)
	r.Env = append(r.Env, argsEnv...)

	output, err := r.CombinedOutput()
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Problem executing cmdi - TestCase: %s", err.Error())))
		return
	}

	c.Ctx.ResponseWriter.Write(output)
}
