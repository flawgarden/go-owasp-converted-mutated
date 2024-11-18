package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01004Controller struct {
	web.Controller
}

func (c *BenchmarkTest01004Controller) Get() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{Name: "BenchmarkTest01004", Value: "bar", Path: c.Ctx.Request.RequestURI, MaxAge: 60 * 3, Secure: true}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "sqli-02/BenchmarkTest01004.html")
}

func (c *BenchmarkTest01004Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	theCookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, theCookie := range theCookies {
		if theCookie.Name == "BenchmarkTest01004" {
			param, _ = url.QueryUnescape(theCookie.Value)
			break
		}
	}

	bar := new(Test).doSomething(c.Ctx.Request, param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

	conn, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	statement, err := conn.Prepare(sqlStr)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	statement.Exec("foo")
	// Add logic to handle results and print them
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	bar := "bobs_your_uncle"
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	}
	return bar
}
