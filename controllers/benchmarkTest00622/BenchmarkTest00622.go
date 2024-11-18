package controllers

import (
	"fmt"
	"os"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00622Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00622Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00622Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00622")

	if param == "" {
		param = ""
	}

	var bar string
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	fileTarget := fmt.Sprintf("%s/%s", "testfiles", bar)
	c.Ctx.Output.Body([]byte(fmt.Sprintf("Access to file: '%s' created.", htmlEscape(fileTarget))))
	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		c.Ctx.Output.Body([]byte(" But file doesn't exist yet."))
	} else {
		c.Ctx.Output.Body([]byte(" And file already exists."))
	}
}

func htmlEscape(str string) string {
	return strings.ReplaceAll(strings.ReplaceAll(str, "&", "&amp;"), "<", "&lt;")
}
