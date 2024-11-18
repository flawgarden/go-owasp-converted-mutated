package controllers

import (
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00995Controller struct {
	web.Controller
}

func (c *BenchmarkTest00995Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{Name: "BenchmarkTest00995", Value: "color", MaxAge: 60 * 3, Secure: true, Path: c.Ctx.Request.RequestURI, Domain: c.Ctx.Request.Host}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)

	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "trustbound-00/BenchmarkTest00995.html")
}

func (c *BenchmarkTest00995Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00995" {
			param = cookie.Value
			break
		}
	}

	bar := new(Test).doSomething(c.Ctx.Request, param)

	c.Ctx.Request.Context().Value("session").(http.ResponseWriter).Header().Set(bar, "10340")

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Item: '%s' with value: '10340' saved in session.", bar)))
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	bar := ""
	if param != "" {
		bar = param
		if spaceIndex := stringIndex(param, ' '); spaceIndex != -1 {
			bar = param[:spaceIndex]
		}
	}
	return bar
}

func stringIndex(s string, r rune) int {
	for i, v := range s {
		if v == r {
			return i
		}
	}
	return -1
}
