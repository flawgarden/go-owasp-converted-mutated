package controllers

import (
	"fmt"
	"os"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00699Controller struct {
	web.Controller
}

func (c *BenchmarkTest00699Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00699Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest00699")
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
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
	case 'C':
	case 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	fileTarget := fmt.Sprintf("%s/Test.txt", strings.TrimSpace(bar))
	c.Ctx.Output.Body([]byte(fmt.Sprintf("Access to file: '%s' created.", fileTarget)))
	if _, err := os.Stat(fileTarget); err == nil {
		c.Ctx.Output.Body([]byte(" And file already exists."))
	} else {
		c.Ctx.Output.Body([]byte(" But file doesn't exist yet."))
	}
}

func main() {
	web.Router("/pathtraver-00/BenchmarkTest00699", &BenchmarkTest00699Controller{})
	web.Run()
}
