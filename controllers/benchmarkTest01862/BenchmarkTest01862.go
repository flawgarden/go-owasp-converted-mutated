package controllers

import (
	"html/template"
	"net/http"
	"net/url"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01862Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01862Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest01862",
		Value:  "whatever",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.URL.Path,
		Domain: c.Ctx.Request.Host,
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "securecookie-00/BenchmarkTest01862.html")
}

func (c *BenchmarkTest01862Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	theCookies := c.Ctx.Request.Cookies()

	param := "noCookieValueSupplied"
	if theCookies != nil {
		for _, theCookie := range theCookies {
			if theCookie.Name == "BenchmarkTest01862" {
				param, _ = url.QueryUnescape(theCookie.Value)
				break
			}
		}
	}

	bar := doSomething(c.Ctx.Request, param)
	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   true,
		HttpOnly: true,
		Path:     c.Ctx.Request.URL.Path,
	}
	http.SetCookie(c.Ctx.ResponseWriter, &cookie)

	c.Ctx.ResponseWriter.Write([]byte("Created cookie: 'SomeCookie': with value: '" + htmlEscape(bar) + "' and secure flag set to: true"))
}

func doSomething(r *http.Request, param string) string {
	return param
}

func htmlEscape(str string) string {
	return template.HTMLEscapeString(str)
}
