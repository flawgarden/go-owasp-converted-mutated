package controllers

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00623Controller struct {
	web.Controller
}

func (c *BenchmarkTest00623Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00623Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00623")
	if param == "" {
		param = ""
	}

	var bar string
	if param != "" {
		decoded, _ := base64.StdEncoding.DecodeString(param)
		bar = string(decoded)
	}

	fileTarget := fmt.Sprintf("%s/Test.txt", bar)
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", fileTarget)))
	if _, err := os.Stat(fileTarget); err == nil {
		c.Ctx.ResponseWriter.Write([]byte(" And file already exists."))
	} else {
		c.Ctx.ResponseWriter.Write([]byte(" But file doesn't exist yet."))
	}
}

func main() {
	web.Router("/pathtraver-00/BenchmarkTest00623", &BenchmarkTest00623Controller{})
	web.Run()
}
