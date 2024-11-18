package controllers

import (
	"net/url"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00322Controller struct {
	web.Controller
}

func (c *BenchmarkTest00322Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00322Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["BenchmarkTest00322"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := "safe!"
	map53101 := map[string]interface{}{
		"keyA-53101": "a_Value",
		"keyB-53101": param,
		"keyC":       "another_Value",
	}
	bar = map53101["keyB-53101"].(string)
	bar = map53101["keyA-53101"].(string)

	c.Ctx.Request.Context().Value("session").(map[string]interface{})["userid"] = bar

	c.Ctx.ResponseWriter.Write([]byte("Item: 'userid' with value: '" + bar + "' saved in session."))
}
