package controllers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

type BenchmarkTest00977Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00977Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest00977",
		Value:  "whatever",
		MaxAge: 180,
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: getDomain(c.Ctx.Request.URL.String()),
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "securecookie-00/BenchmarkTest00977.html")
}

func (c *BenchmarkTest00977Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00977" {
			param, _ = url.QueryUnescape(cookie.Value)
			break
		}
	}

	bar := new(Test).doSomething(c.Ctx.Request, param)

	var output string
	if bar == "" {
		bar = "No cookie value supplied"
	}
	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		HttpOnly: true,
		Path:     c.Ctx.Request.RequestURI,
	}
	http.SetCookie(c.Ctx.ResponseWriter, &cookie)
	output = fmt.Sprintf("Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: false", bar)

	c.Ctx.ResponseWriter.Write([]byte(output))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := param
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	}
	return bar
}

func getDomain(urlStr string) string {
	u, err := url.Parse(urlStr)
	if err != nil {
		return ""
	}
	return u.Hostname()
}
