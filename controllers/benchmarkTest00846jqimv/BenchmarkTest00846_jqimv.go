//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: []
//-------------
//Semgrep analysis results: [89]
//Gosec analysis results: [89, 703]
//CodeQL analysis results: [563]
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00846/BenchmarkTest00846.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/pointers.tmt with name pointer_init2_positive 
//Used extensions: 
//Program:
package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/url"
	"strings"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type SqlInjectionTestController struct {
	web.Controller
}

func (c *SqlInjectionTestController) Get() {
	queryString := c.Ctx.Request.URL.RawQuery

var ptr123 *string = new(string)
var ptr234 *string = new(string)
*ptr234 = queryString
ptr123 = ptr234
queryString = *ptr123

	paramVal := "BenchmarkTest00846="
	paramLoc := -1
	if queryString != "" {
		paramLoc = findParameter(queryString, paramVal)
	}

	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest00846")))
		return
	}

	param := extractParameterValue(queryString, paramLoc, paramVal)
	bar := param

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM users WHERE username='foo' AND password='%s'", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func findParameter(queryString, paramVal string) int {
	return strings.Index(queryString, paramVal)
}

func extractParameterValue(queryString string, paramLoc int, paramVal string) string {
	param := queryString[paramLoc+len(paramVal):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramVal) : ampersandLoc+paramLoc]
	}
	param, _ = url.QueryUnescape(param)
	return param
}

func addSuffix(s *string, suf string) {
	*s = *s + suf
}

func addSuffixDoublePointer(s **string, suf *string) {
	**s = **s + *suf
}

func addSuffixDoublePointerBroken(s **string, suf *string) {
	*s = new(string)
	**s = **s + *suf
}

func swapStrings(a, b *string) {
	temp := *a
	*a = *b
	*b = temp
}

func removeSpaces(s *string) {
    *s = strings.ReplaceAll(*s, " ", "")
}


