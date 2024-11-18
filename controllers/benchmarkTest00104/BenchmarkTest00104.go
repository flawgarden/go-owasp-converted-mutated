package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

type BenchmarkTest00104 struct {
	beego.Controller
}

func (c *BenchmarkTest00104) Get() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
		Name:   "BenchmarkTest00104",
		Value:  "bar",
		Path:   c.Ctx.Request.RequestURI,
		MaxAge: 60 * 3,
		Secure: true,
		Domain: c.Ctx.Request.Host,
	})
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "sqli-00/BenchmarkTest00104.html")
}

func (c *BenchmarkTest00104) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	var param string = "noCookieValueSupplied"
	cookies := c.Ctx.Request.Cookies()

	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00104" {
			decodedValue, _ := url.QueryUnescape(cookie.Value)
			param = decodedValue
			break
		}
	}

	bar := param
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	}

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	_, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.WriteString("Error processing request.")
		return
	}

	c.Ctx.WriteString(fmt.Sprintf("No results can be displayed for query: %s<br> because the Spring batchUpdate method doesn't return results.", sqlStr))
}
