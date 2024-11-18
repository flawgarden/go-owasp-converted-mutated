package controllers

import (
	"fmt"
	"html"
	"net/http"
	"os"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02121Controller struct {
	web.Controller
}

func (c *BenchmarkTest02121Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02121Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02121")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	// Действия с хешированием
	hashValue := hashString(bar)

	fileTarget := "passwordFile.txt"
	f, err := openFile(fileTarget, true)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Could not open file", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("hash_value=%s\n", hashValue))
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Could not write to file", http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", htmlEscape(bar))))
	c.Ctx.ResponseWriter.Write([]byte("Hash Test executed"))
}

func doSomething(param string) string {
	if param != "" {
		return string(param) // Пример обработки параметра
	}
	return ""
}

func hashString(input string) string {
	// Пример хеширования (замените на нужную реализацию)
	return input // Замена на реальное значение хеширования
}

func openFile(fileName string, append bool) (*os.File, error) {
	if append {
		return os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	}
	return os.Create(fileName)
}

func htmlEscape(s string) string {
	return html.EscapeString(s)
}
