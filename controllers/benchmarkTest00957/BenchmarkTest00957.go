package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00957Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00957Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest00957",
		Value:  "FileName",
		MaxAge: 180,
		Secure: true,
		Path:   c.Ctx.Request.URL.Path,
		Domain: c.Ctx.Request.Host,
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "./pathtraver-01/BenchmarkTest00957.html")
}

func (c *BenchmarkTest00957Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string = "noCookieValueSupplied"
	if cookies := c.Ctx.Request.Cookies(); cookies != nil {
		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest00957" {
				param = cookie.Value
				break
			}
		}
	}

	bar := doSomething(param)

	fileName := filepath.Join("TESTFILES_DIR", bar)
	file, err := os.Open(fileName)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Couldn't open InputStream on file: '" + fileName + "'"))
		return
	}
	defer file.Close()

	b := make([]byte, 1000)
	size, err := file.Read(b)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Problem getting InputStream: " + err.Error()))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("The beginning of file: '" + fileName + "' is:\n\n"))
	c.Ctx.ResponseWriter.Write(b[:size])
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = append(valuesList[:0], valuesList[1:]...)
		bar = valuesList[0]
	}
	return bar
}
