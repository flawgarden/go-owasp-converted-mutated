package controllers

import (
	"net/http"
	"net/url"
)

type BenchmarkTest01922 struct{}

func (b *BenchmarkTest01922) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := ""

	if referrer := r.Header.Get("Referer"); referrer != "" {
		param = referrer

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
        Color: param,
    },
}

param = car.Specs.Color

	}

	decodedParam, _ := url.QueryUnescape(param)
	bar := doSomething(decodedParam)

	w.Header().Set("X-XSS-Protection", "0")
	_, _ = w.Write([]byte(bar))
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
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


