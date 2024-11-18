package controllers

import (
	"fmt"
	"net/http"

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

type BenchmarkTest00963 struct {
	beego.Controller
}

func (c *BenchmarkTest00963) Get() {
	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
		Name:   "BenchmarkTest00963",
		Value:  "someSecret",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: c.Ctx.Request.Host,
	})
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "hash-01/BenchmarkTest00963.html")
}

func (c *BenchmarkTest00963) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00963" {
			param = cookie.Value
			break
		}
	}

	bar := new(Test).doSomething(param)

	// Example placeholder for hash processing
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value '%s' processed\n", bar)))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	return param // Simple return; add encoding as necessary
}
