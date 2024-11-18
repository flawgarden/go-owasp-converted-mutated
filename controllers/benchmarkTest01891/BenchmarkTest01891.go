package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	// Initialization code can be added here if required
}

type BenchmarkTest01891Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01891Controller) Get() {
	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
		Name:   "BenchmarkTest01891",
		Value:  "bar",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Input.URL(),
		Domain: getDomain(c.Ctx.Input.URL()),
	})

	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "sqli-04/BenchmarkTest01891.html")
}

func (c *BenchmarkTest01891Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"

	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01891" {
			param, _ = url.QueryUnescape(cookie.Value)
			break
		}
	}

	bar := doSomething(param)
	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", bar)

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

	c.Ctx.ResponseWriter.Write([]byte("Update complete"))
}

func doSomething(param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
}

func getDomain(urlStr string) string {
	parsedUrl, _ := url.Parse(urlStr)
	return parsedUrl.Host
}
