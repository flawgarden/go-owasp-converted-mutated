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

type BenchmarkTest01890 struct {
	web.Controller
}

func (c *BenchmarkTest01890) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{Name: "BenchmarkTest01890", Value: "bar", MaxAge: 60 * 3, Secure: true, Path: c.Ctx.Request.RequestURI, Domain: c.Ctx.Request.Host})

	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "sqli-04/BenchmarkTest01890.html")
}

func (c *BenchmarkTest01890) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	cookies := c.Ctx.Request.Cookies()

	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01890" {
			decodedValue, _ := url.QueryUnescape(cookie.Value)
			param = decodedValue
			break
		}
	}

	bar := doSomething(c.Ctx.Request, param)
	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("Update complete"))
}

func doSomething(r *http.Request, param string) string {
	bar := param
	return bar
}
