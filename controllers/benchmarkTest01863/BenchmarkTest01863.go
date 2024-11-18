package controllers

import (
	"html/template"
	"net/http"
	"net/url"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01863Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01863Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest01863",
		Value:  "whatever",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: getDomain(c.Ctx.Request.URL.String()),
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)

	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "securecookie-00/BenchmarkTest01863.html")
}

func (c *BenchmarkTest01863Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := "noCookieValueSupplied"
	cookies := c.Ctx.Request.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01863" {
			param, _ = url.QueryUnescape(cookie.Value)
			break
		}
	}

	bar := doSomething(param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   true,
		HttpOnly: true,
		Path:     c.Ctx.Request.RequestURI,
	}
	http.SetCookie(c.Ctx.ResponseWriter, &cookie)

	c.Ctx.ResponseWriter.Write([]byte("Created cookie: 'SomeCookie': with value: '" + bar + "' and secure flag set to: true"))
}

func doSomething(param string) string {
	return htmlEscape(param)
}

func htmlEscape(s string) string {
	return template.HTMLEscapeString(s)
}

func getDomain(urlStr string) string {
	u, _ := url.Parse(urlStr)
	return u.Host
}
