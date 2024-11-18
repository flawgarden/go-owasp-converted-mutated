package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00946 struct {
	web.Controller
}

func (c *BenchmarkTest00946) Get() {
	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
		Name:   "BenchmarkTest00946",
		Value:  "someSecret",
		MaxAge: 180, // Store cookie for 3 minutes
		Secure: true,
		Path:   c.Ctx.Request.URL.Path,
		Domain: c.Ctx.Request.Host,
	})

	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "crypto-01/BenchmarkTest00946.html")
}

func (c *BenchmarkTest00946) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string = "noCookieValueSupplied"
	cookies := c.Ctx.Request.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00946" {
			param = cookie.Value
			break
		}
	}

	bar := doSomething(c.Ctx.Request, param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(request *http.Request, param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}
