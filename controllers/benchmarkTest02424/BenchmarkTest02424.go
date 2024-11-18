package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02424Controller struct {
	web.Controller
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (c *BenchmarkTest02424Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest02424Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest02424Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02424")
	if param == "" {
		param = ""
	}


	rememberMeKey := fmt.Sprintf("%d", rand.Int())

	user := "Ingrid"
	testCaseNumber := "02424"
	user += testCaseNumber

	cookieName := "rememberMe" + testCaseNumber

	foundUser := false
	cookies := c.Ctx.Request.Cookies()
	for _, cookie := range cookies {
		if cookieName == cookie.Name {
			if cookie.Value == c.Ctx.Request.Header.Get(cookieName) {
				foundUser = true
			}
		}
	}

	if foundUser {
		c.Ctx.WriteString("Welcome back: " + user + "<br/>")
	} else {
		rememberMe := &http.Cookie{
			Name:     cookieName,
			Value:    rememberMeKey,
			Secure:   true,
			HttpOnly: true,
			Path:     c.Ctx.Request.RequestURI,
		}
		http.SetCookie(c.Ctx.ResponseWriter, rememberMe)
		c.Ctx.Request.Header.Set(cookieName, rememberMeKey)
		c.Ctx.WriteString(user + " has been remembered with cookie: " + rememberMe.Name + " whose value is: " + rememberMe.Value + "<br/>")
	}

	c.Ctx.WriteString("Weak Randomness Test rand.Int() executed")
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		bar = string([]byte(param)) // Placeholder for the actual decoding logic
	}
	return bar
}
