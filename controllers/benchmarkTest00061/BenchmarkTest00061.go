package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00061Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00061Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest00061",
		Value:  "FileName",
		MaxAge: 60 * 3, // Store cookie for 3 minutes
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: c.Ctx.Request.URL.Host,
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "pathtraver-00/BenchmarkTest00061.html")
}

func (c *BenchmarkTest00061Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	theCookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"

	for _, theCookie := range theCookies {
		if theCookie.Name == "BenchmarkTest00061" {
			param = theCookie.Value
			break
		}
	}

	bar := ""
	if param != "" {
		decodedParam, _ := url.PathUnescape(param)
		bar = string([]byte(decodedParam))
	}

	fileTarget := fmt.Sprintf("%s/Test.txt", strings.TrimSpace(bar))
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", fileTarget)))

	if _, err := os.Stat(fileTarget); err == nil {
		c.Ctx.ResponseWriter.Write([]byte(" And file already exists."))
	} else {
		c.Ctx.ResponseWriter.Write([]byte(" But file doesn't exist yet."))
	}
}
