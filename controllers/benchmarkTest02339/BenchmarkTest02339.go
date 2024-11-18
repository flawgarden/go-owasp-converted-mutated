package controllers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02339Controller struct {
	web.Controller
}

func (c *BenchmarkTest02339Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02339Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true

	names := c.Ctx.Request.PostForm
	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest02339" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := doSomething(param)

	cookie := &http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   false,
		HttpOnly: true,
		Path:     c.Ctx.Request.RequestURI,
	}
	http.SetCookie(c.Ctx.ResponseWriter, cookie)

	c.Ctx.Output.Body([]byte("Created cookie: 'SomeCookie': with value: '" + bar + "' and secure flag set to: false"))
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = append(valuesList[:0], valuesList[1:]...) // remove the 1st safe value
		bar = valuesList[1]                                    // get the last 'safe' value
	}
	return bar
}
