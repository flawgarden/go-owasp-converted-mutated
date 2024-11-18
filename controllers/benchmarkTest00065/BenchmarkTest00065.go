package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00065Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00065Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest00065",
		Value:  "FileName",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: strings.Split(c.Ctx.Request.Host, ":")[0],
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)

	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "/pathtraver-00/BenchmarkTest00065.html")
}

func (c *BenchmarkTest00065Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	cookies := c.Ctx.Request.Cookies()

	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00065" {
			param = cookie.Value
			break
		}
	}

	bar := ""
	if param != "" {
		bar = string(param)
	}

	fileName := filepath.Join("testfiles", bar)
	var content string
	file, err := os.Open(fileName)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Couldn't open InputStream on file: '" + fileName + "'"))
		return
	}
	defer file.Close()

	buffer := make([]byte, 1000)
	size, err := file.Read(buffer)
	if err == nil {
		content = string(buffer[:size])
	}

	c.Ctx.ResponseWriter.Write([]byte("The beginning of file: '" + fileName + "' is:\n\n" + content))
}
