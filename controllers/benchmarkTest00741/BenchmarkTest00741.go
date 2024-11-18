package controllers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	// Initialize database connection
}

type BenchmarkTest00741Controller struct {
	web.Controller
}

func (c *BenchmarkTest00741Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00741Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00741")

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}

	cmd := ""
	if os := os.Getenv("OS"); os != "" && strings.Contains(os, "Windows") {
		cmd = "echo "
	}

	argsEnv := []string{"Foo=bar"}
	r := exec.Command(cmd+bar, argsEnv...)
	process, err := r.CombinedOutput()
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Problem executing cmdi - TestCase: %s", err.Error())))
		return
	}
	c.Ctx.ResponseWriter.Write(process)
}
