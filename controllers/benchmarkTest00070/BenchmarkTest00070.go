package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"html"
	"net/http"
	"net/url"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00070Controller struct {
	web.Controller
}

func (c *BenchmarkTest00070Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest00070",
		Value:  "someSecret",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.URL.Path,
		Domain: c.Ctx.Request.URL.Hostname(),
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "hash-00/BenchmarkTest00070.html")
}

func (c *BenchmarkTest00070Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	theCookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, theCookie := range theCookies {
		if theCookie.Name == "BenchmarkTest00070" {
			param, _ = url.QueryUnescape(theCookie.Value)
			break
		}
	}

	bar := param
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	}

	hasher := sha1.New()
	hasher.Write([]byte(bar))
	result := hasher.Sum(nil)
	fileTarget, _ := os.Create("passwordFile.txt")
	defer fileTarget.Close()
	fileTarget.WriteString("hash_value=" + base64.StdEncoding.EncodeToString(result) + "\n")

	c.Ctx.ResponseWriter.Write([]byte("Sensitive value '" + html.EscapeString(bar) + "' hashed and stored<br/>"))
	c.Ctx.ResponseWriter.Write([]byte("Hash Test executed"))
}
