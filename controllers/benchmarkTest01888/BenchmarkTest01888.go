package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

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

type BenchmarkTest01888Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01888Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
		Name:   "BenchmarkTest01888",
		Value:  "bar",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: getHostFromURL(c.Ctx.Request.URL.String()),
	})
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "./sqli-04/BenchmarkTest01888.html")
}

func (c *BenchmarkTest01888Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string = "noCookieValueSupplied"
	if cookies := c.Ctx.Request.Cookies(); cookies != nil {
		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest01888" {
				param, _ = url.QueryUnescape(cookie.Value)
				break
			}
		}
	}

	bar := doSomething(c.Ctx.Request, param)

	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	statement, err := db.Exec(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	_, err = statement.RowsAffected()
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
}

func doSomething(r *http.Request, param string) string {
	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}

func getHostFromURL(u string) string {
	parsedURL, _ := url.Parse(u)
	return parsedURL.Host
}
