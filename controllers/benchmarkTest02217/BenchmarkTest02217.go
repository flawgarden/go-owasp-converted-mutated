package controllers

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02217Controller struct {
	web.Controller
}

func (c *BenchmarkTest02217Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02217Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02217")
	bar := doSomething(param)

	hasher := md5.New()
	input := []byte(bar)
	hasher.Write(input)
	result := hasher.Sum(nil)

	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, fmt.Sprintf("Error opening file: %v", err), http.StatusInternalServerError)
		return
	}
	defer fw.Close()

	if _, err := fw.WriteString(fmt.Sprintf("hash_value=%x\n", result)); err != nil {
		http.Error(c.Ctx.ResponseWriter, fmt.Sprintf("Error writing to file: %v", err), http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", input)))
	c.Ctx.ResponseWriter.Write([]byte("Hash Test executed<br/>"))
}

func doSomething(param string) string {
	var bar string
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
