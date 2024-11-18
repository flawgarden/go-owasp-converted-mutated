package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type HashTestController struct {
	web.Controller
}

func (c *HashTestController) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
		Name:   "BenchmarkTest01848",
		Value:  "someSecret",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.URL.Path,
		Domain: c.Ctx.Request.Host,
	})
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "hash-02/BenchmarkTest01848.html")
}

func (c *HashTestController) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01848" {
			param = cookie.Value
			break
		}
	}

	bar := doSomething(param)

	hashValue := hashInput(bar)

	fileTarget := "passwordFile.txt"
	ioutil.WriteFile(fileTarget, []byte(fmt.Sprintf("hash_value=%x\n", hashValue)), 0644)

	output := fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", bar)
	c.Ctx.ResponseWriter.Write([]byte(output))
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // Remove the 1st safe value
		bar = valuesList[0]         // Get the last 'safe' value
	}
	return bar
}

func hashInput(input string) []byte {
	// Hashing logic here (omitted for brevity)
	// This function should hash the input and return the hash value as bytes
	return []byte{}
}
