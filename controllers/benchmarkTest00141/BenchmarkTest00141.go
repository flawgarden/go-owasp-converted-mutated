package controllers

import (
	"crypto/sha1"
	"fmt"
	"net/url"
	"os"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00141Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00141Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00141Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	param := c.Ctx.Request.Header.Get("BenchmarkTest00141")
	param, _ = url.QueryUnescape(param)

	bar := ""
	num := 106

	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

	h := sha1.New()
	input := []byte(bar)
	h.Write(input)
	result := h.Sum(nil)

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		c.Ctx.Output.Body([]byte("Unable to open file"))
		return
	}
	defer file.Close()

	if _, err := file.WriteString("hash_value=" + fmt.Sprintf("%x", result) + "\n"); err != nil {
		c.Ctx.Output.Body([]byte("Unable to write to file"))
		return
	}

	c.Ctx.Output.Body([]byte("Sensitive value '" + param + "' hashed and stored<br/>"))
}
