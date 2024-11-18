package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

type BenchmarkTest struct {
	beego.Controller
}

func (c *BenchmarkTest) Get() {
	c.doPost(c.Ctx.ResponseWriter, c.Ctx.Request)
}

func (c *BenchmarkTest) doPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := request.URL.Query().Get("BenchmarkTest01518")
	if param == "" {
		param = ""
	}


	rand := float32(0) // Replace with your own random number generation method
	rememberMeKey := fmt.Sprintf("%f", rand)[2:]

	user := "Floyd"
	fullClassName := fmt.Sprintf("%T", c)
	testCaseNumber := fullClassName[strings.LastIndex(fullClassName, ".")+1+len("BenchmarkTest"):]

	user += testCaseNumber

	cookieName := "rememberMe" + testCaseNumber

	cookies := request.Cookies()
	foundUser := false
	for _, cookie := range cookies {
		if cookieName == cookie.Name {
			if cookie.Value == request.Context().Value(cookieName) {
				foundUser = true
			}
		}
	}

	if foundUser {
		response.Write([]byte("Welcome back: " + user + "<br/>"))
	} else {
		http.SetCookie(response, &http.Cookie{
			Name:     cookieName,
			Value:    rememberMeKey,
			Secure:   true,
			HttpOnly: true,
			Path:     request.URL.Path,
		})
		request.Context().Value(cookieName)
		response.Write([]byte(user + " has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}

	response.Write([]byte("Weak Randomness Test executed"))
}

func (c *BenchmarkTest) doSomething(request *http.Request, param string) string {
	a := param
	b := a + " SafeStuff"
	b = b[:len(b)-len("Chars")] + "Chars"
	results := map[string]interface{}{"key": b}
	c1 := results["key"].(string)
	c2 := c1[:len(c1)-1]
	f := c2[:len(c2)-1] // Custom decode logic: replace with actual decoding if needed
	bar := f            // Do something meaningful here
	return bar
}
