package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"html"
	"net/http"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00872Controller struct {
	web.Controller
}

func (c *BenchmarkTest00872Controller) Get() {
	c.DoPost(c.Ctx.ResponseWriter, c.Ctx.Request)
}

func (c *BenchmarkTest00872Controller) Post() {
	c.DoPost(c.Ctx.ResponseWriter, c.Ctx.Request)
}

func (c *BenchmarkTest00872Controller) DoPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest00872")

	bar := ""
	if param != "" {
		bar = param
	}

	hash := sha1.New()
	hash.Write([]byte(bar))
	result := hash.Sum(nil)

	fileTarget := "passwordFile.txt"
	fw, _ := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer fw.Close()
	fw.WriteString("hash_value=" + base64.StdEncoding.EncodeToString(result) + "\n")

	output, _ := json.Marshal(map[string]string{
		"message": "Sensitive value '" + html.EscapeString(bar) + "' hashed and stored",
	})

	w.Write(output)
}
