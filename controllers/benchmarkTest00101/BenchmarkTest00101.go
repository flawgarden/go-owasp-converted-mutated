package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00101Controller struct {
	web.Controller
}

func (c *BenchmarkTest00101Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	cookie := http.Cookie{Name: "BenchmarkTest00101", Value: "bar", MaxAge: 60 * 3, Secure: true, Path: c.Ctx.Request.RequestURI, Domain: c.Ctx.Request.URL.Host}
	http.SetCookie(c.Ctx.ResponseWriter, &cookie)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "sqli-00/BenchmarkTest00101.html")
}

func (c *BenchmarkTest00101Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	theCookies := c.Ctx.Request.Cookies()

	param := "noCookieValueSupplied"
	for _, theCookie := range theCookies {
		if theCookie.Name == "BenchmarkTest00101" {
			param = theCookie.Value
			break
		}
	}

	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME=? AND PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	_, err = statement.Exec("foo")
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}

	// Здесь можно добавить код для отображения результата
}
