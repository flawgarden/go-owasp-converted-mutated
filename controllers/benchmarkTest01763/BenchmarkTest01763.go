package controllers

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01763Controller struct {
	web.Controller
}

func (c *BenchmarkTest01763Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01763Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest01763")
	bar := new(Test).doSomething(c.Ctx.Request, param)

	hash := sha256.New()
	if _, err := hash.Write([]byte(bar)); err != nil {
		panic(err)
	}
	result := hash.Sum(nil)

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("hash_value=%x\n", result)); err != nil {
		panic(err)
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", bar)))
	c.Ctx.ResponseWriter.Write([]byte("Hash Test executed"))
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	thing := createThing()
	return thing.doSomething(param)
}

type ThingInterface interface {
	doSomething(param string) string
}

func createThing() ThingInterface {
	return &Thing{}
}

type Thing struct{}

func (t *Thing) doSomething(param string) string {
	return param
}
