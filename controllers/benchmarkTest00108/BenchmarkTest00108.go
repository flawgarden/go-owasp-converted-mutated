package controllers

import (
	"database/sql"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	// Database initialization should be here
}

type BenchmarkTest00108 struct {
	web.Controller
}

func (c *BenchmarkTest00108) Get() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	c.Ctx.SetCookie("BenchmarkTest00108", "bar", 60*3)
	c.Redirect("/sqli-00/BenchmarkTest00108.html", 302)
}

func (c *BenchmarkTest00108) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	cookieValue := "noCookieValueSupplied"
	cookies := c.Ctx.Request.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00108" {
			cookieValue = cookie.Value
			break
		}
	}

	bar := determineBar(cookieValue)

	sqlStatement := "SELECT * from USERS where USERNAME='foo' and PASSWORD='" + bar + "'"

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStatement)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}

	// Further processing as needed
}

func determineBar(param string) string {
	guess := "ABC"
	switchTarget := guess[2]

	var bar string
	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}
	return bar
}
