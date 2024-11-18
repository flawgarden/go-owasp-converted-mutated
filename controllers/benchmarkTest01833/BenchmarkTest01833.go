package controllers

import (
	"net/http"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01833Controller struct {
	web.Controller
}

func (c *BenchmarkTest01833Controller) Get() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest01833",
		Value:  "FileName",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: c.Ctx.Request.Host,
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "pathtraver-02/BenchmarkTest01833.html")
}

func (c *BenchmarkTest01833Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	var param string = "noCookieValueSupplied"
	cookies := c.Ctx.Request.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01833" {
			param = cookie.Value
			break
		}
	}

	bar := doSomething(c.Ctx.Request, param)

	fileTarget := "TESTFILES_DIR/" + bar
	c.Ctx.Output.Body([]byte("Access to file: '" + fileTarget + "' created."))

	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		c.Ctx.Output.Body([]byte(" But file doesn't exist yet."))
	} else {
		c.Ctx.Output.Body([]byte(" And file already exists."))
	}
}

func doSomething(request *http.Request, param string) string {
	var bar string = "safe!"
	map9325 := make(map[string]interface{})
	map9325["keyA-9325"] = "a-Value"
	map9325["keyB-9325"] = param
	map9325["keyC"] = "another-Value"
	bar = map9325["keyB-9325"].(string)

	return bar
}
