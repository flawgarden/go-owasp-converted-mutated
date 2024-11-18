package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type CookieHandler struct {
	web.Controller
}

func (c *CookieHandler) Get() {
	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
		Name:   "BenchmarkTest01845",
		Value:  "someSecret",
		MaxAge: 180, // Store cookie for 3 minutes
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: c.Ctx.Request.Host,
	})
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "hash-02/BenchmarkTest01845.html")
}

func (c *CookieHandler) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string = "noCookieValueSupplied"
	if cookies := c.Ctx.Request.Cookies(); cookies != nil {
		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest01845" {
				param = cookie.Value
				break
			}
		}
	}

	bar := doSomething(c.Ctx.Request, param)
	id := strings.TrimSpace(bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where id='%s'", id)
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

func doSomething(r *http.Request, param string) string {
	bar := "noModification"
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	}
	return bar
}
