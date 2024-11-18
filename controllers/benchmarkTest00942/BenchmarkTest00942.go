package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

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

type BenchmarkTest00942 struct {
	beego.Controller
}

func (c *BenchmarkTest00942) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest00942",
		Value:  "someSecret",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
	}
	c.Ctx.ResponseWriter.Header().Add("Set-Cookie", userCookie.String())

	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "crypto-01/BenchmarkTest00942.html")
}

func (c *BenchmarkTest00942) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00942" {
			param, _ = url.QueryUnescape(cookie.Value)
			break
		}
	}

	bar := newTest().doSomething(c.Ctx.Request, param)

	rand := time.Now().UnixNano()
	iv := make([]byte, 16)
	for i := range iv {
		iv[i] = byte(rand % 256)
		rand /= 256
	}

	// encryption logic omitted for simplicity
	result := []byte("encrypted_data") // Replace with actual encryption

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("secret_value=%s\n", string(result))); err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(c.Ctx.ResponseWriter, "Sensitive value: '%s' encrypted and stored<br/>", bar)
}

type test struct{}

func newTest() *test {
	return &test{}
}

func (t *test) doSomething(request *http.Request, param string) string {
	if param != "" {
		return string(param) // simplistic decode simulation
	}
	return ""
}
