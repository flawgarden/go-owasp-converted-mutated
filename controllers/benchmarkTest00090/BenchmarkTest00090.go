package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00090Controller struct {
	web.Controller
}

func (c *BenchmarkTest00090Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest00090",
		Value:  "ls",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: c.Ctx.Request.Host,
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "cmdi-00/BenchmarkTest00090.html")
}

func (c *BenchmarkTest00090Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := "noCookieValueSupplied"
	cookies := c.Ctx.Request.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00090" {
			param, _ = url.QueryUnescape(cookie.Value)
			break
		}
	}

	bar := ""
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	cmd := ""
	if strings.Contains(strings.ToLower(c.Ctx.Request.UserAgent()), "windows") {
		cmd = "cmd /c echo "
	}

	out, err := exec.Command(cmd + bar).Output()
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Problem executing cmdi - TestCase: %s", err.Error())))
		return
	}

	c.Ctx.ResponseWriter.Write(out)
}
