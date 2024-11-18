package controllers

import (
	"fmt"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02015 struct {
	web.Controller
}

func (c *BenchmarkTest02015) Get() {
	c.Post()
}

func (c *BenchmarkTest02015) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	for name := range c.Ctx.Request.Header {
		if isCommonHeader(name) {
			continue
		}
		param = name
		break
	}

	bar := doSomething(param)

	c.Ctx.Request.Header.Set(bar, "10340")
	fmt.Fprintf(c.Ctx.ResponseWriter, "Item: '%s' with value: 10340 saved in session.", encodeForHTML(bar))
}

func isCommonHeader(header string) bool {
	commonHeaders := []string{"Content-Type", "Content-Length", "Host"}
	for _, h := range commonHeaders {
		if strings.EqualFold(h, header) {
			return true
		}
	}
	return false
}

func doSomething(param string) string {
	var bar string
	if param != "" {
		bar = string([]byte(param)) // Simulating Base64 decode then encode
	}
	return bar
}

func encodeForHTML(input string) string {
	// Simple HTML encoding simulation
	return strings.ReplaceAll(input, "<", "&lt;")
}

func main() {
	web.Router("/trustbound-01/BenchmarkTest02015", &BenchmarkTest02015{})
	web.Run()
}
