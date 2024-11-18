package controllers

import (
	"database/sql"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01885Controller struct {
	web.Controller
}

func (c *BenchmarkTest01885Controller) Get() {
	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
		Name:   "BenchmarkTest01885",
		Value:  "bar",
		Path:   c.Ctx.Input.URI(),
		MaxAge: 60 * 3,
		Secure: true,
	})

	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "sqli-04/BenchmarkTest01885.html")
}

func (c *BenchmarkTest01885Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01885" {
			param = cookie.Value
			break
		}
	}

	bar := doSomething(param)

	sqlStr := "SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='" + bar + "'"

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

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}
	return bar
}
