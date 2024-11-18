package controllers

import (
	"net/http"
	"net/url"
	"text/template"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01875Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01875Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest01875",
		Value:  "color",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: getDomain(c.Ctx.Request.URL),
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	c.TplName = "trustbound-01/BenchmarkTest01875.html"
}

func (c *BenchmarkTest01875Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	theCookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, theCookie := range theCookies {
		if theCookie.Name == "BenchmarkTest01875" {
			param, _ = url.QueryUnescape(theCookie.Value)
			break
		}
	}

	bar := doSomething(param)
	c.Ctx.Request.Context().Value("session").(http.ResponseWriter).Write([]byte("Item: '" + bar + "' with value: '10340' saved in session."))
}

func doSomething(param string) string {
	return htmlEscape(param)
}

func htmlEscape(s string) string {
	return template.HTMLEscapeString(s)
}

func getDomain(u *url.URL) string {
	return u.Hostname()
}
