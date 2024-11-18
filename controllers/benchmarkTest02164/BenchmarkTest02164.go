package controllers

import (
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02164Controller struct {
	web.Controller
}

func (c *BenchmarkTest02164Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02164Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02164")
	if param == "" {
		param = ""
	}


	// SecureRandom emulation
	rememberMeKey := fmt.Sprintf("%d", GetSecureRandomLong())
	user := "SafeLogan"
	fullClassName := "BenchmarkTest02164"
	testCaseNumber := fullClassName[len("BenchmarkTest"):len(fullClassName)]

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
			Secure:   true,
			HttpOnly: true,
			Path:     c.Ctx.Request.RequestURI,
		})
		c.Ctx.Request.Context().Value(cookieName) // Set session attribute
		c.Ctx.ResponseWriter.Write([]byte(user + " has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}
	c.Ctx.ResponseWriter.Write([]byte("Weak Randomness Test executed"))
}

func doSomething(r *http.Request, param string) string {
	return param // Simple implementation, sanitize as required
}

func GetSecureRandomLong() int64 {
	// Simplified function to return a random long value
	return 1234567890 // For demonstration purposes, replace with actual secure random implementation
}
