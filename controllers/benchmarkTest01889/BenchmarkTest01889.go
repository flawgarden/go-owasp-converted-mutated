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

type BenchmarkTest01889 struct {
	web.Controller
}

func (c *BenchmarkTest01889) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{Name: "BenchmarkTest01889", Value: "bar", Path: c.Ctx.Input.URI(), MaxAge: 60 * 3, Secure: true}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "/sqli-04/BenchmarkTest01889.html")
}

func (c *BenchmarkTest01889) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	theCookies := c.Ctx.Request.Cookies()

	param := "noCookieValueSupplied"
	if theCookies != nil {
		for _, cookie := range theCookies {
			if cookie.Name == "BenchmarkTest01889" {
				param, _ = url.QueryUnescape(cookie.Value)
				break
			}
		}
	}

	bar := doSomething(c.Ctx.Request, param)

	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("Update complete."))
}

func doSomething(r *http.Request, param string) string {
	return param
}
