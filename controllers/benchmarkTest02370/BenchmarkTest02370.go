package controllers

import (
	"os"
	"strings"

	"github.com/beego/beego/v2/server/web"
	"golang.org/x/net/html"
)

type BenchmarkTest02370Controller struct {
	web.Controller
}

func (c *BenchmarkTest02370Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02370Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	names := c.Ctx.Request.URL.Query()
	for name, values := range names {
		for _, value := range values {
			if value == "BenchmarkTest02370" {
				param = name
				break
			}
		}
	}

	bar := doSomething(param)

	file, err := os.Open("employees.xml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		panic(err)
	}

	expression := "//Employee[@emplid='" + bar + "']"
	nodeList := selectNodes(doc, expression)

	c.Ctx.ResponseWriter.Write([]byte("Your query results are: <br/>"))
	for _, node := range nodeList {
		c.Ctx.ResponseWriter.Write([]byte(node.Data + "<br/>"))
	}
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

func selectNodes(doc *html.Node, expression string) []*html.Node {
	// Пример функции для выбора узлов из XML-документа.
	var nodes []*html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "Employee" {
			for _, attr := range n.Attr {
				if attr.Key == "emplid" && attr.Val == strings.TrimPrefix(expression, "//Employee[@emplid='") {
					nodes = append(nodes, n)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return nodes
}
