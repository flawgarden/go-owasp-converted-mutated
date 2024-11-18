package controllers

import (
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00002Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00002Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest00002",
		Value:  "FileName",
		MaxAge: 180, // Store cookie for 3 minutes
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: c.Ctx.Request.Host,
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "pathtraver-00/BenchmarkTest00002.html")
}

func (c *BenchmarkTest00002Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := c.Ctx.Request.Cookies()

	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00002" {
			param, _ = url.QueryUnescape(cookie.Value)
			break
		}
	}

	var fileName string
	var file *os.File

	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	fileName = filepath.Join("path/to/testfiles", param)

	file, _ = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	c.Ctx.ResponseWriter.Write([]byte("Now ready to write to file: " + fileName))
}
