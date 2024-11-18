package controllers

import (
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02384Controller struct {
	web.Controller
}

func (c *BenchmarkTest02384Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02384Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02384")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	base := "ou=users,ou=system"
	filter := "(&(objectclass=person)(uid=" + bar + "))"

	// Simulate LDAP Searching (pseudo code)
	results := ldapSearch(base, filter)

	if len(results) > 0 {
		for _, result := range results {
			c.Ctx.Output.Body([]byte("LDAP query results:<br>" +
				"Record found with name " + result["uid"] + "<br>" +
				"Address: " + result["street"] + "<br>"))
		}
	} else {
		c.Ctx.Output.Body([]byte("LDAP query results: nothing found for query: " + htmlEncode(filter)))
	}
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:]

		bar = valuesList[1]
	}
	return bar
}

func ldapSearch(base, filter string) []map[string]string {
	// Pseudo LDAP search function
	var results []map[string]string
	// Simulated data retrieval
	if strings.Contains(filter, "someuid") {
		results = append(results, map[string]string{
			"uid":    "someuid",
			"street": "123 Main St",
		})
	}
	return results
}

func htmlEncode(input string) string {
	// Simple HTML encode for demo purposes
	return strings.ReplaceAll(input, "&", "&amp;")
}
