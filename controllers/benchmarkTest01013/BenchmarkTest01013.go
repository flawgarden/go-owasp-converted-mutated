package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01013 struct {
	web.Controller
}

func (c *BenchmarkTest01013) Get() {
	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
		Name:   "BenchmarkTest01013",
		Value:  "2222",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.URL.Path,
	})

	c.Ctx.Output.Body([]byte("Cookie set"))
}

func (c *BenchmarkTest01013) Post() {
	theCookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range theCookies {
		if cookie.Name == "BenchmarkTest01013" {
			param = cookie.Value
			break
		}
	}

	bar := new(Test).doSomething(c.Ctx.Request, param)

	xmlFile := "employees.xml"
	// Assume readXML method reads the XML file and returns a parsed document
	xmlDocument, err := readXML(xmlFile)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error parsing XML", http.StatusInternalServerError)
		return
	}

	expression := fmt.Sprintf("/Employees/Employee[@emplid='%s']", bar)
	nodeList := evaluateXPath(xmlDocument, expression)

	var results []string
	for _, node := range nodeList {
		results = append(results, node.TextContent())
	}

	c.Ctx.Output.Body([]byte(strings.Join(results, "<br/>")))
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	bar := param
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	}
	return bar
}

func readXML(file string) (interface{}, error) {
	// Your XML reading logic here
	return nil, nil
}

func evaluateXPath(doc interface{}, expression string) []Node {
	// Your XPath evaluation logic here
	return nil
}

type Node struct {
	// Your Node structure here
	TextContent func() string
}
