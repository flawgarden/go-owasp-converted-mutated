package controllers

import (
	"net/http"
	"net/url"
	"os"

	"github.com/beego/beego/v2/server/web"
	"golang.org/x/net/html"
)

type BenchmarkTest01224Controller struct {
	web.Controller
}

func (c *BenchmarkTest01224Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01224Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01224Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Request.Header.Get("BenchmarkTest01224")
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(c.Ctx.Request, param)

	file, err := os.Open("employees.xml")
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error parsing HTML", http.StatusInternalServerError)
		return
	}

	expression := "//Employee[@emplid='" + bar + "']"
	result := evaluateXPath(doc, expression)

	c.Ctx.ResponseWriter.Write([]byte("Your query results are: " + result + "<br/>"))
}

func evaluateXPath(doc *html.Node, expression string) string {
	// Implement XPath evaluation logic here
	return "mocked result"
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}

	return bar
}
