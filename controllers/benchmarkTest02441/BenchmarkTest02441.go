package controllers

import (
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02441 struct {
	web.Controller
}

func (c *BenchmarkTest02441) Get() {
	c.Post()
}

func (c *BenchmarkTest02441) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02441")
	if param == "" {
		param = ""
	}


	stuff := getSecureRandom()
	rememberMeKey := fmt.Sprintf("%.4f", stuff)[2:]

	user := "SafeGayle"
	testCaseNumber := "02441"
	user += testCaseNumber

	cookieName := "rememberMe" + testCaseNumber
	cookies := c.Ctx.Request.Cookies()

	foundUser := false
	for _, cookie := range cookies {
		if cookieName == cookie.Name {
			if cookie.Value == c.Ctx.Request.URL.Query().Get(cookieName) {
				foundUser = true
			}
		}
	}

	if foundUser {
		c.Ctx.ResponseWriter.Write([]byte("Welcome back: " + user + "<br/>"))
	} else {
		http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
			Name:     cookieName,
			Value:    rememberMeKey,
			Secure:   true,
			HttpOnly: true,
			Path:     c.Ctx.Request.RequestURI,
		})
		c.Ctx.Request.URL.Query().Set(cookieName, rememberMeKey)
		c.Ctx.ResponseWriter.Write([]byte(user + " has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}

	c.Ctx.ResponseWriter.Write([]byte("Weak Randomness Test executed"))
}

func doSomething(param string) string {
	bar := ""
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func getSecureRandom() float64 {
	return 0.0 // Use actual secure random logic
}
