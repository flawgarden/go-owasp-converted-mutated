package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00421Controller struct {
	web.Controller
}

func (c *BenchmarkTest00421Controller) Get() {
	c.ServeJSON()
}

func (c *BenchmarkTest00421Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00421")
	if param == "" {
		param = ""
	}


	rand.Seed(time.Now().UnixNano())
	randNumber := rand.Intn(99)
	rememberMeKey := fmt.Sprintf("%d", randNumber)

	user := "SafeInga"
	fullClassName := fmt.Sprintf("%T", c)
	testCaseNumber := fullClassName[len("controllers.BenchmarkTest"):]

	user += testCaseNumber
	cookieName := "rememberMe" + testCaseNumber

	foundUser := false
	cookies := c.Ctx.Request.Cookies()
	for _, cookie := range cookies {
		if cookieName == cookie.Name {
			if cookie.Value == c.GetSession(cookieName) {
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
		c.SetSession(cookieName, rememberMeKey)
		c.Ctx.ResponseWriter.Write([]byte(user + " has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}
	c.Ctx.ResponseWriter.Write([]byte("Weak Randomness Test executed"))
}
