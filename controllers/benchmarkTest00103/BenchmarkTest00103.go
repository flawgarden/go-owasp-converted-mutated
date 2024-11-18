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

type BenchmarkTest00103 struct {
	web.Controller
}

func (c *BenchmarkTest00103) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{Name: "BenchmarkTest00103", Value: "bar", MaxAge: 60 * 3, Secure: true, Path: c.Ctx.Request.RequestURI, Domain: c.Ctx.Request.Host}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)

	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "sqli-00/BenchmarkTest00103.html")
}

func (c *BenchmarkTest00103) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	theCookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, theCookie := range theCookies {
		if theCookie.Name == "BenchmarkTest00103" {
			param = theCookie.Value
			break
		}
	}

	bar := param
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	}

	sqlStr := fmt.Sprintf("SELECT TOP 1 USERNAME from USERS where USERNAME='foo' and PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var results string
	err = db.QueryRow(sqlStr).Scan(&results)
	if err == sql.ErrNoRows {
		c.Ctx.ResponseWriter.Write([]byte("No results returned for query: " + htmlEscape(sqlStr)))
	} else if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
	} else {
		c.Ctx.ResponseWriter.Write([]byte("Your results are: " + htmlEscape(results)))
	}
}

func htmlEscape(input string) string {
	return url.QueryEscape(input)
}
