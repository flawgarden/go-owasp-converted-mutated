package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01828Controller struct {
	web.Controller
}

func init() {
	web.Router("/crypto-02/BenchmarkTest01828", &BenchmarkTest01828Controller{})
}

func (c *BenchmarkTest01828Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
		Name:   "BenchmarkTest01828",
		Value:  "someSecret",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.RequestURI,
		Domain: c.Ctx.Request.Host,
	})

	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "/crypto-02/BenchmarkTest01828.html")
}

func (c *BenchmarkTest01828Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01828" {
			param = cookie.Value
			break
		}
	}

	bar := doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	input := []byte(bar)
	if _, err = db.Exec("INSERT INTO sensitive_data (value) VALUES (?)", input); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Database insertion error", http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value: '%s' stored", bar)))
}

func doSomething(param string) string {
	return param + "_SafeStuff"
}
