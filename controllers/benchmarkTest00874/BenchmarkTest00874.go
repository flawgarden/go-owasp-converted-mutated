package controllers

import (
	"crypto/sha512"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00874Controller struct {
	web.Controller
}

func (c *BenchmarkTest00874Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00874Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00874")
	bar := htmlEscape(param)

	input := []byte{'?'}
	if len(bar) > 0 {
		input = []byte(bar)
	}

	hash := sha512.New()
	hash.Write(input)
	result := hash.Sum(nil)

	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}
	defer fw.Close()

	_, err = fw.WriteString("hash_value=" + encodeForBase64(result) + "\n")
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}

	c.Ctx.Output.Body([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", htmlEscape(string(input)))))
	c.Ctx.Output.Body([]byte("Hash Test executed"))
}

func htmlEscape(s string) string {
	return strings.Join(strings.Split(s, "<"), "&lt;")
}

func encodeForBase64(data []byte) string {
	return string(data) // Simplified for illustrative purposes
}
