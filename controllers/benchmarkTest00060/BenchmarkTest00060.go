package controllers

import (
	"net/http"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00060Controller struct {
	web.Controller
}

func (c *BenchmarkTest00060Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest00060",
		Value:  "FileName",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.URL.Path,
		Domain: c.Ctx.Request.Host,
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	c.Ctx.Output.Body([]byte("Cookie set"))
}

func (c *BenchmarkTest00060Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	theCookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"

	for _, theCookie := range theCookies {
		if theCookie.Name == "BenchmarkTest00060" {
			param = theCookie.Value
			break
		}
	}

	bar := ""
	if param != "" {
		bar = string(param) // Simulating base64 decode for simplicity
	}

	fileTarget := "./" + bar
	c.Ctx.Output.Body([]byte("Access to file: '" + fileTarget + "' created."))

	if _, err := os.Stat(fileTarget); err == nil {
		c.Ctx.Output.Body([]byte(" And file already exists."))
	} else {
		c.Ctx.Output.Body([]byte(" But file doesn't exist yet."))
	}
}
