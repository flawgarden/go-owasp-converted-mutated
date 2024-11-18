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

type BenchmarkTest00966 struct {
	beego.Controller
}

func (c *BenchmarkTest00966) Get() {
	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
		Name:     "BenchmarkTest00966",
		Value:    "someSecret",
		MaxAge:   180,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		Path:     c.Ctx.Input.URI(),
		Domain:   c.Ctx.Input.Host(),
	})
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "./hash-01/BenchmarkTest00966.html")
}

func (c *BenchmarkTest00966) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00966" {
			param = cookie.Value
			break
		}
	}

	bar := c.doSomething(param)

	// Implement hashing logic here (SHA5 etc.) as necessary

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Ctx.Output.Body([]byte(fmt.Sprintf("Sensitive value '%s' processed<br/>", bar)))
}

func (c *BenchmarkTest00966) doSomething(param string) string {
	bar := param
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	}
	return bar
}
