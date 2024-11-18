package controllers

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	// registry database
}

type BenchmarkTest01706Controller struct {
	web.Controller
}

func (c *BenchmarkTest01706Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01706Controller) Post() {
	queryString := c.Ctx.Request.URL.RawQuery
	paramval := "BenchmarkTest01706="
	paramLoc := -1

	if queryString != "" {
		paramLoc = findParamLocation(queryString, paramval)
	}

	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest01706' in query string."))
		return
	}

	param := extractParam(queryString, paramLoc, paramval)
	param, _ = url.QueryUnescape(param)


	r := generateRandomInt()
	rememberMeKey := strconv.Itoa(r)
	user := "SafeIngrid"
	testCaseNumber := "01706"
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
		c.Ctx.ResponseWriter.Write([]byte("Welcome back: " + user + "<br/>"))
	} else {
		http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
			Name:     cookieName,
			Value:    rememberMeKey,
			Secure:   true,
			HttpOnly: true,
			Path:     c.Ctx.Request.RequestURI,
		})
		c.Ctx.ResponseWriter.Write([]byte(user + " has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}

	c.Ctx.ResponseWriter.Write([]byte("Weak Randomness Test executed"))
}

func findParamLocation(queryString, paramval string) int {
	return 0 // реализация для поиска местоположения параметра
}

func extractParam(queryString string, paramLoc int, paramval string) string {
	return "" // реализация для извлечения параметра
}

func doSomething(param string) string {
	return param
}

func generateRandomInt() int {
	return 0 // реализация генерации случайного числа
}
