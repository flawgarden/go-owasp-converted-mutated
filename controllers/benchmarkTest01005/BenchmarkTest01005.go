package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	// Initialize the database connection here if necessary
}

type BenchmarkTest01005Controller struct {
	web.Controller
}

func (c *BenchmarkTest01005Controller) Get() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest01005",
		Value:  "bar",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: getDomain(c.Ctx.Request.URL.String()),
	}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "html/sqli-02/BenchmarkTest01005.html")
}

func (c *BenchmarkTest01005Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	cookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01005" {
			param, _ = url.QueryUnescape(cookie.Value)
			break
		}
	}

	bar := new(innerClass).doSomething(c.Ctx.Request, param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("No results can be displayed for query: %s<br>", sqlStr)))
}

type innerClass struct{}

func (t *innerClass) doSomething(request *http.Request, param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		if len(valuesList) > 0 {
			bar = valuesList[0] // get the param value
		}
	}
	return bar
}

func getDomain(requestURL string) string {
	u, err := url.Parse(requestURL)
	if err != nil {
		return ""
	}
	return u.Host
}
