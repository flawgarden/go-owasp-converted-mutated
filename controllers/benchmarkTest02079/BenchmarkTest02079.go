package controllers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02079Controller struct {
	web.Controller
}

func (c *BenchmarkTest02079Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02079Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["BenchmarkTest02079"]
	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)


	// Simulate the random key generation
	stuff := "generatedKey" // Placeholder for random key
	rememberMeKey := strings.TrimPrefix(stuff, "0.")

	user := "SafeGayle"
	testCaseNumber := "02079"

	user += testCaseNumber
	cookieName := "rememberMe" + testCaseNumber

	cookies := c.Ctx.Request.Cookies()
	foundUser := false
	for _, cookie := range cookies {
		if cookie.Name == cookieName && cookie.Value == c.Ctx.Request.Context().Value(cookieName) {
			foundUser = true
			break
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
		c.Ctx.Request.Context().Value(cookieName + "_Value")
		c.Ctx.ResponseWriter.Write([]byte(user + " has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}
	c.Ctx.ResponseWriter.Write([]byte("Weak Randomness Test executed"))
}

func doSomething(request *http.Request, param string) string {
	// Simulate some operation with the parameter
	return param
}
