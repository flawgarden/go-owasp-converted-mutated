package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00970Controller struct {
	web.Controller
}

func (c *BenchmarkTest00970Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest00970",
		Value:  "ECHOOO",
		MaxAge: 60 * 3, // Store cookie for 3 minutes
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: getDomain(c.Ctx.Request.URL.String()),
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "cmdi-01/BenchmarkTest00970.html")
}

func (c *BenchmarkTest00970Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	theCookies := c.Ctx.Request.Cookies()

	param := "noCookieValueSupplied"
	for _, theCookie := range theCookies {
		if theCookie.Name == "BenchmarkTest00970" {
			param, _ = url.QueryUnescape(theCookie.Value)
			break
		}
	}

	bar := new(Test).doSomething(c.Ctx.Request, param)

	argList := []string{}
	if isWindows() {
		argList = append(argList, "cmd.exe", "/c")
	} else {
		argList = append(argList, "sh", "-c")
	}
	argList = append(argList, "echo "+bar)

	cmd := exec.Command(argList[0], argList[1:]...)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Problem executing cmdi - exec.Command Test Case")
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(out)
}

func getDomain(requestURL string) string {
	parsedURL, _ := url.Parse(requestURL)
	return parsedURL.Host
}

func isWindows() bool {
	return runtime.GOOS == "windows"
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	bar := ""
	num := 106

	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}
