package controllers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00565Controller struct {
	web.Controller
}

func (c *BenchmarkTest00565Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00565Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := c.Ctx.Request.PostForm

	for name := range names {
		values := c.Ctx.Request.PostForm[name]
		if values != nil {
			for _, value := range values {
				if value == "BenchmarkTest00565" {
					param = name
					flag = false
					break
				}
			}
		}
		if !flag {
			break
		}
	}

	bar := EncodeForHTML(param)

	var responseMessage string
	if param == "" {
		responseMessage = "No cookie value supplied"
	} else {
		cookie := http.Cookie{
			Name:     "SomeCookie",
			Value:    bar,
			Secure:   false,
			HttpOnly: true,
			Path:     c.Ctx.Request.RequestURI,
		}
		http.SetCookie(c.Ctx.ResponseWriter, &cookie)

		responseMessage = "Created cookie: 'SomeCookie': with value: '" + EncodeForHTML(bar) + "' and secure flag set to: false"
	}

	c.Ctx.ResponseWriter.Write([]byte(responseMessage))
}

func EncodeForHTML(input string) string {
	// HTML encoding implementation
	return input // Replace with actual implementation
}
