package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00077Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00077Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest00077",
		Value:  "ECHOOO",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Input.URL(),
		Domain: c.Ctx.Request.Host,
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "cmdi-00/BenchmarkTest00077.html")
}

func (c *BenchmarkTest00077Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	theCookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, theCookie := range theCookies {
		if theCookie.Name == "BenchmarkTest00077" {
			param, _ = url.QueryUnescape(theCookie.Value)
			break
		}
	}

	bar := ""
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

	argList := []string{}
	if isWindows() {
		argList = append(argList, "cmd.exe", "/c")
	} else {
		argList = append(argList, "sh", "-c")
	}
	argList = append(argList, "echo "+bar)

	cmd := exec.Command(argList[0], argList[1:]...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Problem executing cmdi - java.lang.ProcessBuilder(java.util.List) Test Case")
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func isWindows() bool {
	return runtime.GOOS == "windows"
}
