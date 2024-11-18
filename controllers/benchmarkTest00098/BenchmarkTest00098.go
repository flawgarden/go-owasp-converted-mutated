package controllers

import (
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00098Controller struct {
	web.Controller
}

func (c *BenchmarkTest00098Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest00098",
		Value:  "my_user_id",
		MaxAge: 60 * 3, // Store cookie for 3 minutes
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: getHostname(c.Ctx.Request.URL.String()),
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "trustbound-00/BenchmarkTest00098.html")
}

func (c *BenchmarkTest00098Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	theCookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, theCookie := range theCookies {
		if theCookie.Name == "BenchmarkTest00098" {
			param, _ = url.QueryUnescape(theCookie.Value)
			break
		}
	}
	bar := escapeHTML(param)
	c.Ctx.Request.Header.Set("userid", bar)
	c.Ctx.ResponseWriter.Write([]byte("Item: 'userid' with value: '" + bar + "' saved in session."))
}

func getHostname(urlStr string) string {
	u, _ := url.Parse(urlStr)
	return u.Hostname()
}

func escapeHTML(input string) string {
	return url.QueryEscape(input)
}
