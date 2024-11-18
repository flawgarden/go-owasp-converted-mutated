package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01150Controller struct {
	web.Controller
}

func (c *BenchmarkTest01150Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01150Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	var param string

	if header := c.Ctx.Input.Header("BenchmarkTest01150"); header != "" {
		param = header
	}

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(param)

	block, err := des.NewCipher(make([]byte, 8))
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Problem executing crypto", http.StatusInternalServerError)
		return
	}

	ciphertext := make([]byte, len(bar))
	mode := cipher.NewCBCEncrypter(block, make([]byte, des.BlockSize))
	mode.CryptBlocks(ciphertext, []byte(bar))

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Cannot open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte(fmt.Sprintf("secret_value=%x\n", ciphertext)))
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Cannot write to file", http.StatusInternalServerError)
		return
	}

	_, _ = c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value: '%s' encrypted and stored<br/>", bar)))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	if param == "" {
		return ""
	}
	valuesList := []string{"safe", param, "moresafe"}
	valuesList = valuesList[1:] // remove the 1st safe value
	return valuesList[0]        // get the param value
}
