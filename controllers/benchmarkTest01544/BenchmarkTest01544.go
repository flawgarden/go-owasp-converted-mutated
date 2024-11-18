package controllers

import (
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01544 struct {
	beego.Controller
}

func (c *BenchmarkTest01544) Get() {
	c.Post()
}

func (c *BenchmarkTest01544) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01544")
	if param == "" {
		param = ""
	}

	rand := float32(0) // Replace with secure random float logic
	rememberMeKey := fmt.Sprintf("%.2f", rand)[2:]

	user := "SafeFloyd"
	fullClassName := "BenchmarkTest01544"
	testCaseNumber := fullClassName[len("BenchmarkTest"):]
	user += testCaseNumber

	cookieName := "rememberMe" + testCaseNumber
	foundUser := false
	cookies := c.Ctx.Request.Cookies()
	for _, cookie := range cookies {
		if cookieName == cookie.Name {
			if cookie.Value == c.Ctx.Request.URL.Query().Get(cookieName) {
				foundUser = true
				break
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
			Path:     c.Ctx.Request.URL.Path,
		})
		c.Ctx.Request.URL.Query().Set(cookieName, rememberMeKey)
		c.Ctx.ResponseWriter.Write([]byte(user + " has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}

	c.Ctx.ResponseWriter.Write([]byte("Weak Randomness Test java.security.SecureRandom.nextFloat() executed"))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = append(valuesList[:0], valuesList[1:]...) // remove the 1st safe value
		bar = valuesList[0]                                    // get the last 'safe' value
	}
	return bar
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

func main() {
	beego.Router("/weakrand-03/BenchmarkTest01544", &BenchmarkTest01544{})
	beego.Run()
}
