package controllers

import (
	"net/http"
	"net/url"
	"text/template"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00088Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00088Controller) Get() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest00088",
		Value:  "whatever",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: getHostFromURL(c.Ctx.Request.URL.String()),
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	c.Ctx.Redirect(http.StatusFound, "/securecookie-00/BenchmarkTest00088.html")
}

func (c *BenchmarkTest00088Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	theCookies := c.Ctx.Request.Cookies()

	param := "noCookieValueSupplied"
	for _, theCookie := range theCookies {
		if theCookie.Name == "BenchmarkTest00088" {
			param = theCookie.Value
			break
		}
	}

	decodedParam, _ := url.QueryUnescape(param)
	bar := htmlEscape(decodedParam)

	var output string
	if bar == "" {
		output = "No cookie value supplied"
	} else {
		cookie := http.Cookie{
			Name:     "SomeCookie",
			Value:    bar,
			Secure:   true,
			HttpOnly: true,
			Path:     c.Ctx.Request.RequestURI,
		}
		http.SetCookie(c.Ctx.ResponseWriter, &cookie)
		output = "Created cookie: 'SomeCookie': with value: '" + htmlEscape(bar) + "' and secure flag set to: true"
	}

	c.Ctx.WriteString(output)
}

func getHostFromURL(requestURL string) string {
	u, _ := url.Parse(requestURL)
	return u.Hostname()
}

func htmlEscape(s string) string {
	return template.HTMLEscapeString(s)
}
