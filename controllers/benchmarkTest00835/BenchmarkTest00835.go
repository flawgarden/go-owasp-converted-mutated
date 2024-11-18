package controllers

import (
	"fmt"
	"html"
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest00835Controller struct {
	http.Handler
}

func (c *BenchmarkTest00835Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.doPost(w, r)
	} else if r.Method == http.MethodPost {
		c.doPost(w, r)
	}
}

func (c *BenchmarkTest00835Controller) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest00835="
	paramLoc := strings.Index(queryString, paramval)

	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest00835"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}

	param, _ = url.QueryUnescape(param)
	bar := "alsosafe"

	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}

	r.Context().Value("session").(http.ResponseWriter).Header().Set("userid", bar)
	fmt.Fprintf(w, "Item: 'userid' with value: '%s' saved in session.", html.EscapeString(bar))
}
