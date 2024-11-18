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

type BenchmarkTest00115 struct {
	web.Controller
}

func (c *BenchmarkTest00115) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest00115",
		Value:  "bar",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: c.Ctx.Request.Host,
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "./sqli-00/BenchmarkTest00115.html")
}

func (c *BenchmarkTest00115) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	theCookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"

	for _, theCookie := range theCookies {
		if theCookie.Name == "BenchmarkTest00115" {
			param, _ = url.QueryUnescape(theCookie.Value)
			break
		}
	}

	bar := "safe!"
	map11928 := make(map[string]interface{})
	map11928["keyA-11928"] = "a-Value"
	map11928["keyB-11928"] = param
	map11928["keyC"] = "another-Value"
	bar = map11928["keyB-11928"].(string)

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
}
