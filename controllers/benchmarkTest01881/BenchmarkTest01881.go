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

type BenchmarkTest01881 struct {
	web.Controller
}

func (c *BenchmarkTest01881) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	userCookie := http.Cookie{Name: "BenchmarkTest01881", Value: "bar", MaxAge: 60 * 3, Secure: true, Path: c.Ctx.Request.RequestURI, Domain: c.Ctx.Request.Host}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)

	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "sqli-04/BenchmarkTest01881.html")
}

func (c *BenchmarkTest01881) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01881" {
			param = cookie.Value
			break
		}
	}

	bar := doSomething(param)

	sqlStatement := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

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

	fmt.Fprintf(c.Ctx.ResponseWriter, "No results can be displayed for query: %s<br>because the query method doesn't return results.", htmlEscape(sqlStatement))
}

func doSomething(param string) string {
	bar := "safe!"
	map20027 := make(map[string]interface{})
	map20027["keyA-20027"] = "a-Value"
	map20027["keyB-20027"] = param
	map20027["keyC"] = "another-Value"
	bar = map20027["keyB-20027"].(string)

	return bar
}

func htmlEscape(s string) string {
	return url.QueryEscape(s)
}
