package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01846Controller struct {
	web.Controller
}

func (c *BenchmarkTest01846Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
		Name:     "BenchmarkTest01846",
		Value:    "someSecret",
		MaxAge:   60 * 3,
		SameSite: http.SameSiteLaxMode,
		Path:     c.Ctx.Request.RequestURI,
	})
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "hash-02/BenchmarkTest01846.html")
}

func (c *BenchmarkTest01846Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	cookies := c.Ctx.Request.Cookies()

	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01846" {
			param = cookie.Value
			break
		}
	}

	bar := doSomething(param)

	// Simulate hashing process
	fmt.Printf("Hashing value: %s\n", bar)

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value '%s' processed<br/>", bar)))
}

func doSomething(param string) string {
	return strings.ReplaceAll(param, "<", "&lt;")
}
