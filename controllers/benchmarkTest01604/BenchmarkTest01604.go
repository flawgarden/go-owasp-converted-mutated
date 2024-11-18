package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01604Controller struct {
	web.Controller
}

func (c *BenchmarkTest01604Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest01604Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest01604Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01604")
	if param == "" {
		param = "No cookie value supplied"
	}

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    param,
		Secure:   true,
		HttpOnly: true,
		Path:     c.Ctx.Request.RequestURI,
	}
	http.SetCookie(c.Ctx.ResponseWriter, &cookie)

	output := map[string]string{
		"message": "Created cookie: 'SomeCookie': with value: '" + param + "' and secure flag set to: true",
	}
	response, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(response)
}
