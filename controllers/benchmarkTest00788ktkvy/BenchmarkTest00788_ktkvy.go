package controllers

import (
	"net/url"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00788Controller struct {
	web.Controller
}

func (c *BenchmarkTest00788Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00788Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Input.URI()
	paramval := "BenchmarkTest00788="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramval)
		if queryString[paramLoc:] != paramval {
			c.Ctx.Output.Body([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest00788' in query string."))
			return
		}
	}

car := struct {
    Make  string
    Model string
    Specs struct {
        Year int
        Color string
    }
}{
    Make:  "Toyota",
    Model: "X5 AMG",
    Specs: struct {
        Year  int
        Color string
    }{
        Year:  2020,
        Color: queryString,
    },
}

queryString = car.Make

	param := queryString[paramLoc+len(paramval):]
	decodedParam, err := url.QueryUnescape(param)
	if err != nil {
		c.Ctx.Output.Body([]byte("Error decoding parameter."))
		return
	}

	fileName := decodedParam
	fos, err := os.Create(fileName)
	if err != nil {
		c.Ctx.Output.Body([]byte("Couldn't open FileOutputStream on file: '" + fileName + "'"))
		return
	}
	defer fos.Close()

	c.Ctx.Output.Body([]byte("Now ready to write to file: " + fileName))
}

func createPoint(x, y string) struct {
    X string
    Y string
} {
    return struct {
        X string
        Y string
    }{
        X: x,
        Y: y,
    }
}


