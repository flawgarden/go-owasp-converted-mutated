package controllers

import (
	"encoding/base64"
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02685Controller struct {
	web.Controller
}

func (c *BenchmarkTest02685Controller) Get() {
	c.post()
}

func (c *BenchmarkTest02685Controller) Post() {
	c.post()
}

func (c *BenchmarkTest02685Controller) post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest02685")
	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	output := fmt.Sprintf("Formatted like: %s and %s.", "a", bar)
	c.Ctx.ResponseWriter.Write([]byte(output))
}

func doSomething(param string) string {
	var bar string
	if param != "" {
		decoded, err := base64.StdEncoding.DecodeString(base64.StdEncoding.EncodeToString([]byte(param)))
		if err == nil {
			bar = string(decoded)
		}
	}
	return bar
}

func main() {
	web.Router("/xss-05/BenchmarkTest02685", &BenchmarkTest02685Controller{})
	web.Run()
}
