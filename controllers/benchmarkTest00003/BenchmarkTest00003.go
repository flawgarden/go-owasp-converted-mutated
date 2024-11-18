package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

type BenchmarkTest00003Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00003Controller) Get() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest00003",
		Value:  "someSecret",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: c.Ctx.Request.Host,
	}
	http.SetCookie(c.Ctx.Output.Context.ResponseWriter, &userCookie)

	http.ServeFile(c.Ctx.Output.Context.ResponseWriter, c.Ctx.Request, "hash-00/BenchmarkTest00003.html")
}

func (c *BenchmarkTest00003Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	theCookies := c.Ctx.Request.Cookies()

	param := "noCookieValueSupplied"
	for _, theCookie := range theCookies {
		if theCookie.Name == "BenchmarkTest00003" {
			param, _ = url.QueryUnescape(theCookie.Value)
			break
		}
	}

	input := []byte("?")
	if len(param) > 0 {
		input = []byte(param)
	}

	hashValue := fmt.Sprintf("%x", input) // Simple hash representation for illustration
	fileTarget := "passwordFile.txt"
	fw, _ := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer fw.Close()
	fw.WriteString(fmt.Sprintf("hash_value=%s\n", hashValue))
	c.Ctx.Output.Context.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", param)))
	c.Ctx.Output.Context.ResponseWriter.Write([]byte("Hash Test executed"))
}
