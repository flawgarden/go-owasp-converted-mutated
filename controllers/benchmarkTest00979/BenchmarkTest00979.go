package controllers

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00979Controller struct {
	web.Controller
}

func (c *BenchmarkTest00979Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
		Name:   "BenchmarkTest00979",
		Value:  ".",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: c.Ctx.Request.Host,
	})
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "cmdi-01/BenchmarkTest00979.html")
}

func (c *BenchmarkTest00979Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00979" {
			param = cookie.Value
			break
		}
	}

	bar := new(test).doSomething(c.Ctx.Request, param)

	var cmd string
	var args []string
	if strings.Contains(strings.ToLower(c.Ctx.Request.UserAgent()), "windows") {
		cmd = "cmd.exe"
		args = []string{"/c", "echo " + bar}
	} else {
		cmd = "sh"
		args = []string{"-c", "ls " + bar}
	}


	cmdOut, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Problem executing cmdi - TestCase: " + err.Error()))
		return
	}
	c.Ctx.ResponseWriter.Write(cmdOut)
}

type test struct{}

func (t *test) doSomething(req *http.Request, param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}
	return bar
}
