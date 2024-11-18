package controllers

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02706Controller struct {
	web.Controller
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (c *BenchmarkTest02706Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest02706Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest02706Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	rememberMeKey := strconv.Itoa(rand.Int())
	user := "Ingrid"
	testCaseNumber := "BenchmarkTest02706"
	user += testCaseNumber
	cookieName := "rememberMe" + testCaseNumber

	var foundUser bool
	cookies := c.Ctx.Request.Cookies()
	for _, cookie := range cookies {
		if cookieName == cookie.Name {
			if cookie.Value == c.GetSession(cookieName).(string) {
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
	c.Ctx.ResponseWriter.Write([]byte("Weak Randomness Test math/rand.Int executed"))
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		bar = param // здесь имитация декодирования Base64
	}
	return bar
}
