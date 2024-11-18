package controllers

import (
	"database/sql"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02700Controller struct {
	web.Controller
}

func init() {
	// Initialize database connection
	_, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
}

func (c *BenchmarkTest02700Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest02700Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest02700Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")


	rememberMeKey := generateRememberMeKey()

	user := "Byron"
	fullClassName := "BenchmarkTest02700"
	user += fullClassName[len("BenchmarkTest"):]

	cookieName := "rememberMe" + fullClassName

	if foundUser, _ := checkCookie(c.Ctx.Request, cookieName); foundUser {
		c.Ctx.ResponseWriter.Write([]byte("Welcome back: " + user + "<br/>"))
	} else {
		http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
			Name:     cookieName,
			Value:    rememberMeKey,
			Secure:   true,
			HttpOnly: true,
		})
		c.Ctx.ResponseWriter.Write([]byte(user + " has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}

	c.Ctx.ResponseWriter.Write([]byte("Weak Randomness Test java.util.Random.nextBytes() executed"))
}

func doSomething(param string) string {
	if param != "" {
		return string(param) // Mocked for simplicity
	}
	return ""
}

func generateRememberMeKey() string {
	// Mocked function, replace with actual implementation
	return "mockedRememberMeKey"
}

func checkCookie(req *http.Request, cookieName string) (bool, error) {
	cookies := req.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == cookieName {
			// Again, mocked for demonstration
			return true, nil
		}
	}
	return false, nil
}
