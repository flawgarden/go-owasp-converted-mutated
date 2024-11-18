package controllers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01951 struct {
	beego.Controller
}

func (c *BenchmarkTest01951) Get() {
	c.Post()
}

func (c *BenchmarkTest01951) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Request.Header.Get("BenchmarkTest01951")
	param, _ = url.QueryUnescape(param)

	rand := generateRandomFloat()
	rememberMeKey := fmt.Sprintf("%f", rand)[2:]

	user := "SafeFloyd"
	testCaseNumber := "01951"
	user += testCaseNumber

	cookieName := "rememberMe" + testCaseNumber

	cookie, err := c.Ctx.Request.Cookie(cookieName)
	foundUser := false
	if err == nil && cookie.Value == c.Ctx.Request.Context().Value(cookieName) {
		foundUser = true
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
		c.Ctx.Request.Context().Value(cookieName) // set in context
		c.Ctx.ResponseWriter.Write([]byte(user + " has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}

	c.Ctx.ResponseWriter.Write([]byte("Weak Randomness Test executed"))
}

func doSomething(param string) string {
	return fmt.Sprintf("%s", param) // Simple escaping for example
}

func generateRandomFloat() float64 {
	return 0.123456 // Dummy value for example
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

func main() {
	beego.Router("/weakrand-04/BenchmarkTest01951", &BenchmarkTest01951{})
	beego.Run()
}
