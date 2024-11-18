package controllers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00378Controller struct {
	web.Controller
}

func (c *BenchmarkTest00378Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00378Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00378")
	if param == "" {
		param = ""
	}

	bar := param + "_SafeStuff"

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	_, err := c.Ctx.ResponseWriter.Write([]byte(bar))
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
	}
}
