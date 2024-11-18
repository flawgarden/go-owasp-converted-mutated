package controllers

import (
	"encoding/base64"
	"html"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00303Controller struct {
	web.Controller
}

func (c *BenchmarkTest00303Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00303Controller) Post() {
	response := c.Ctx.ResponseWriter
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Request.Header.Get("BenchmarkTest00303")
	param, _ = url.QueryUnescape(param)

	var bar string
	if param != "" {
		bar = string(base64.StdEncoding.EncodeToString([]byte(param)))
	}

	cmd := ""
	a1 := ""
	a2 := ""
	args := []string{}

	if strings.Contains(os.Getenv("OS"), "Windows") {
		a1 = "cmd.exe"
		a2 = "/c"
		cmd = "echo"
		args = []string{a1, a2, cmd + " " + bar}
	} else {
		a1 = "sh"
		a2 = "-c"
		cmd = "ping -c1 "
		args = []string{a1, a2, cmd + bar}
	}

	r := exec.Command(args[0], args[1:]...)
	out, err := r.Output()
	if err != nil {
		http.Error(response, html.EscapeString(err.Error()), http.StatusInternalServerError)
		return
	}

	response.Write(out)
}
