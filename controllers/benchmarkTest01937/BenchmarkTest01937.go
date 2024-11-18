package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01937Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01937Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01937Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01937Controller) DoPost() {
	response := c.Ctx.ResponseWriter
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if header := c.Ctx.Request.Header.Get("BenchmarkTest01937"); header != "" {
		param = header
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(c.Ctx.Request, param)

	cmd := ""
	if strings.Contains(strings.ToLower(os.Getenv("OS")), "windows") {
		cmd = "echo "
	}

	if out, err := exec.Command(cmd + bar).CombinedOutput(); err != nil {
		fmt.Println("Problem executing cmdi - TestCase")
		response.Write([]byte(err.Error()))
		return
	} else {
		response.Write(out)
	}
}

func doSomething(request *http.Request, param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}
	return bar
}
