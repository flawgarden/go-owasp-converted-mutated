package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"math/rand"
	"net/http"
	"time"

	"github.com/beego/beego/v2/server/web"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type SqlInjectionVuln1Controller struct {
	web.Controller
}

func (c *SqlInjectionVuln1Controller) Get() {
	id := c.GetString("id")
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where id=%s", id)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *SqlInjectionVuln1Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02159")
	if param == "" {
		param = ""
	}


	rand.Seed(time.Now().UnixNano())
	rememberMeKey := fmt.Sprintf("%f", rand.Float64())[2:]

	user := "SafeDonna"
	className := fmt.Sprintf("%T", c)
	testCaseNumber := className[len("controllers."):]

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
		c.Ctx.Request.Context().Value(cookieName).(map[string]interface{})[cookieName] = rememberMeKey
		c.Ctx.ResponseWriter.Write([]byte(user + " has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}
	c.Ctx.ResponseWriter.Write([]byte("Weak Randomness Test executed"))
}

func doSomething(param string) string {
	a66015 := param
	b66015 := a66015 + " SafeStuff"
	b66015 = b66015[:len(b66015)-1] + "Chars"
	map66015 := map[string]interface{}{"key66015": b66015}
	c66015 := map66015["key66015"].(string)
	d66015 := c66015[:len(c66015)-1]
	e66015 := d66015
	f66015 := e66015

	return f66015
}
