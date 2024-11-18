package controllers

import (
	"net/http"
	"net/url"
	"os/exec"
	"runtime"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01852Controller struct {
	web.Controller
}

func (c *BenchmarkTest01852Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest01852",
		Value:  "ECHOOO",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.URL.Path,
		Domain: mustGetDomain(c.Ctx.Request.URL.String()),
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "./cmdi-02/BenchmarkTest01852.html")
}

func (c *BenchmarkTest01852Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, theCookie := range cookies {
		if theCookie.Name == "BenchmarkTest01852" {
			param, _ = url.QueryUnescape(theCookie.Value)
			break
		}
	}

	bar := doSomething(param)

	var args []string
	if isWindows() {
		args = []string{"cmd.exe", "/c", "echo " + bar}
	} else {
		args = []string{"sh", "-c", "echo " + bar}
	}

	cmd := execCommand(args)
	if err := cmd.Start(); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Problem executing command", http.StatusInternalServerError)
		return
	}

	cmd.Wait()
}

func doSomething(param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
}

func mustGetDomain(urlStr string) string {
	parsedURL, _ := url.Parse(urlStr)
	return parsedURL.Host
}

func isWindows() bool {
	return runtime.GOOS == "windows"
}

func execCommand(args []string) *exec.Cmd {
	return exec.Command(args[0], args[1:]...)
}
