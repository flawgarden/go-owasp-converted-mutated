package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00059 struct {
	web.Controller
}

func (c *BenchmarkTest00059) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest00059",
		Value:  "someSecret",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Input.URL(),
		Domain: c.Ctx.Input.Host(),
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "crypto-00/BenchmarkTest00059.html")
}

func (c *BenchmarkTest00059) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00059" {
			param, _ = url.QueryUnescape(cookie.Value)
			break
		}
	}

	bar := param + "_SafeStuff"

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.WriteString("Database connection error")
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.WriteString("Database query error")
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.WriteString("JSON marshaling error")
		return
	}

	c.Ctx.ResponseWriter.Write(output)
}
