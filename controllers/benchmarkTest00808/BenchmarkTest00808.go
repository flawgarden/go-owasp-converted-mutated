package controllers

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00808Controller struct {
	web.Controller
}

func (c *BenchmarkTest00808Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest00808Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest00808Controller) doPost() {
	response := c.Ctx.ResponseWriter
	request := c.Ctx.Request
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := request.URL.RawQuery
	paramVal := "BenchmarkTest00808="
	paramLoc := -1

	if queryString != "" {
		paramLoc = findParamLoc(queryString, paramVal)
	}

	if paramLoc == -1 {
		response.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest00808' in query string."))
		return
	}

	param := extractParamValue(queryString, paramLoc, paramVal)
	param = decodeParam(param)

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}

	response.Header().Set("X-XSS-Protection", "0")
	// Print bar (unsafe operation)
	fmt.Fprintf(response, bar)
}

func findParamLoc(queryString, paramVal string) int {
	return strings.Index(queryString, paramVal)
}

func extractParamValue(queryString string, paramLoc int, paramVal string) string {
	param := queryString[paramLoc+len(paramVal):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramVal) : paramLoc+ampersandLoc]
	}
	return param
}

func decodeParam(param string) string {
	decodedParam, err := url.QueryUnescape(param)
	if err != nil {
		return ""
	}
	return decodedParam
}
