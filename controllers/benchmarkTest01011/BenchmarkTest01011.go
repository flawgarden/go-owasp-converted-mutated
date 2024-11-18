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

type BenchmarkTest01011Controller struct {
	web.Controller
}

func (c *BenchmarkTest01011Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	cookie := &http.Cookie{Name: "BenchmarkTest01011", Value: "bar", MaxAge: 60 * 3, Secure: true, Path: c.Ctx.Request.RequestURI}
	http.SetCookie(c.Ctx.ResponseWriter, cookie)

	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "sqli-02/BenchmarkTest01011.html")
}

func (c *BenchmarkTest01011Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01011" {
			param, _ = url.QueryUnescape(cookie.Value)
			break
		}
	}

	bar := new(Test).doSomething(c.Ctx.Request, param)
	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
}

type Test struct{}

func (t *Test) doSomething(req *http.Request, param string) string {
	bar := ""
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}
