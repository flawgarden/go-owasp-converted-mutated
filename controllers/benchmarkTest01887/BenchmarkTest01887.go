package controllers

import (
	"database/sql"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01887Controller struct {
	web.Controller
}

func (c *BenchmarkTest01887Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
		Name:   "BenchmarkTest01887",
		Value:  "bar",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: c.Ctx.Request.Host,
	})
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "sqli-04/BenchmarkTest01887.html")
}

func (c *BenchmarkTest01887Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01887" {
			param = cookie.Value
			break
		}
	}

	bar := doSomething(c.Ctx.Request, param)

	sqlStr := "INSERT INTO users (username, password) VALUES ('foo','" + bar + "')"

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
}

func doSomething(request *http.Request, param string) string {
	return param
}
