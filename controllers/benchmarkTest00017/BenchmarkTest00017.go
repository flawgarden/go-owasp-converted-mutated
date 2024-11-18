package controllers

import (
	"html"
	"net/url"
	"os/exec"
	"runtime"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00017Controller struct {
	web.Controller
}

func (c *BenchmarkTest00017Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00017Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["BenchmarkTest00017"]
	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	cmd := ""
	osName := runtime.GOOS
	if strings.Contains(osName, "windows") {
		cmd = "cmd /c echo "
	}

	output, err := exec.Command(cmd + param).CombinedOutput()
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte(html.EscapeString(err.Error())))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
