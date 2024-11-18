package controllers

import (
	"database/sql"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00051Controller struct {
	web.Controller
}

func init() {
	sql.Open("mysql", source)
}

func (c *BenchmarkTest00051Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00051")

	var a1, a2 string
	if os := runtime.GOOS; os == "windows" {
		a1 = "cmd.exe"
		a2 = "/c"
	} else {
		a1 = "sh"
		a2 = "-c"
	}
	args := []string{a1, a2, "echo " + param}

	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.Output.Body([]byte("Error executing command"))
		return
	}
	c.Ctx.Output.Body(output)
}
