package controllers

import (
	"crypto/sha512"
	"io/ioutil"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02678Controller struct {
	web.Controller
}

func (c *BenchmarkTest02678Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02678Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.Ctx.Input.Query("BenchmarkTest02678")
	bar := doSomething(param)

	hash := sha512.New()
	hash.Write([]byte(bar))
	result := hash.Sum(nil)

	file, err := ioutil.TempFile("", "passwordFile.txt")
	if err != nil {
		c.Ctx.WriteString("Error creating file")
		return
	}
	defer file.Close()
	file.WriteString("hash_value=" + string(result) + "\n")

	c.Ctx.WriteString("Sensitive value '" + param + "' hashed and stored<br/>")
	c.Ctx.WriteString("Hash Test sha512 executed")
}

func doSomething(param string) string {
	if param == "" {
		return ""
	}
	return string([]byte{param[0]})
}
