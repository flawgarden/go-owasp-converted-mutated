package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"
	"strconv"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest struct {
	web.Controller
}

func (c *BenchmarkTest) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
		Name:   "BenchmarkTest00968",
		Value:  "ECHOOO",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: c.Ctx.Request.Host,
	})
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "cmdi-01/BenchmarkTest00968.html")
}

func (c *BenchmarkTest) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00968" {
			param = cookie.Value
			break
		}
	}

	bar := param

	id, err := strconv.Atoi(bar)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%d", id)
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
