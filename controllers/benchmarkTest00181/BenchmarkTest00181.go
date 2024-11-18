package controllers

import (
	"fmt"
	"html"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00181Controller struct {
	web.Controller
}

func (c *BenchmarkTest00181Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00181Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Request.Header.Get("BenchmarkTest00181")
	param, _ = url.QueryUnescape(param)


	rand.Seed(time.Now().UnixNano())
	rnd := rand.Float64()

	rememberMeKey := fmt.Sprintf("%f", rnd)[2:]

	user := "SafeDonna"
	fullClassName := "BenchmarkTest00181Controller"
	testCaseNumber := fullClassName[len("controllers.BenchmarkTest"):]

	user += testCaseNumber
	cookieName := "rememberMe" + testCaseNumber

	foundUser := false
	cookies := c.Ctx.Request.Cookies()
	for _, cookie := range cookies {
		if cookieName == cookie.Name {
			if cookie.Value == c.Ctx.Request.Context().Value(cookieName) {
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
			Path:     c.Ctx.Request.RequestURI,
			Secure:   true,
			HttpOnly: true,
		})
		c.Ctx.Request.Context().Value(cookieName) // emulating session attribute
		c.Ctx.ResponseWriter.Write([]byte(user + " has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}

	c.Ctx.ResponseWriter.Write([]byte("Weak Randomness Test executed"))
}

func htmlEscape(s string) string {
	return string(html.EscapeString(s))
}
