package controllers

import (
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02212 struct {
	web.Controller
}

func (c *BenchmarkTest02212) Get() {
	c.Post()
}

func (c *BenchmarkTest02212) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02212")
	bar := doSomething(param)

	hashValue, err := hash(bar)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	fileTarget := "passwordFile.txt"
	err = appendToFile(fileTarget, hashValue)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", bar)))
	c.Ctx.ResponseWriter.Write([]byte("Hash Test executed"))
}

func doSomething(param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
}

func hash(input string) (string, error) {
	// Implement hashing logic here (e.g. using SHA1)
	// Return the hash value
	return "hash_value_placeholder", nil
}

func appendToFile(filePath string, data string) error {
	// Implement file appending logic here
	return nil
}
