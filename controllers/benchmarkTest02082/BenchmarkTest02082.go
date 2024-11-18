package controllers

import (
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02082Controller struct {
	web.Controller
}

func (c *BenchmarkTest02082Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02082Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	headers := c.Ctx.Request.Header["BenchmarkTest02082"]
	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	rand.Seed(time.Now().UnixNano())
	r := rand.Int()
	rememberMeKey := strconv.Itoa(r)

	user := "SafeIngrid"
	fullClassName := "BenchmarkTest02082"
	testCaseNumber := fullClassName[len("BenchmarkTest"):]
	user += testCaseNumber

	cookieName := "rememberMe" + testCaseNumber

	foundUser := false
	cookies := c.Ctx.Request.Cookies()
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
		rememberMe := http.Cookie{
			Name:     cookieName,
			Value:    rememberMeKey,
			Secure:   true,
			HttpOnly: true,
			Path:     c.Ctx.Request.RequestURI,
		}
		http.SetCookie(c.Ctx.ResponseWriter, &rememberMe)
		c.Ctx.Request.URL.Query().Set(cookieName, rememberMeKey)
		c.Ctx.ResponseWriter.Write([]byte(user + " has been remembered with cookie: " + rememberMe.Name + " whose value is: " + rememberMe.Value + "<br/>"))
	}

	c.Ctx.ResponseWriter.Write([]byte("Weak Randomness Test rand.Int() executed"))
}

func doSomething(param string) string {
	bar := param
	return bar
}
