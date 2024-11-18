package controllers

import (
	"html"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00958Controller struct {
	web.Controller
}

func (c *BenchmarkTest00958Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest00958",
		Value:  "FileName",
		MaxAge: 180, // Store cookie for 3 minutes
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: c.Ctx.Request.Host,
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "/pathtraver-01/BenchmarkTest00958.html")
}

func (c *BenchmarkTest00958Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	theCookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, theCookie := range theCookies {
		if theCookie.Name == "BenchmarkTest00958" {
			param, _ = url.QueryUnescape(theCookie.Value)
			break
		}
	}

	bar := Test{}.doSomething(c.Ctx.Request, param)

	fileName := filepath.Join(os.Getenv("TESTFILES_DIR"), bar)
	file, err := os.Open(fileName)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Couldn't open InputStream on file: '" + fileName + "'"))
		return
	}
	defer file.Close()

	b := make([]byte, 1000)
	size, _ := file.Read(b)
	c.Ctx.ResponseWriter.Write([]byte("The beginning of file: '" + html.EscapeString(fileName) + "' is:\n\n"))
	c.Ctx.ResponseWriter.Write([]byte(html.EscapeString(string(b[:size]))))
}

type Test struct{}

func (t Test) doSomething(request *http.Request, param string) string {
	bar := "safe!"
	map37053 := map[string]interface{}{
		"keyA-37053": "a_Value",
		"keyB-37053": param,
		"keyC":       "another_Value",
	}
	bar = map37053["keyB-37053"].(string)
	bar = map37053["keyA-37053"].(string)
	return bar
}
