package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest00073Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest00073Controller) Get() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := &http.Cookie{Name: "BenchmarkTest00073", Value: "someSecret", MaxAge: 60 * 3, Secure: true, Path: c.Request.URL.Path, Domain: c.Request.URL.Host}
	http.SetCookie(c.ResponseWriter, userCookie)

	http.ServeFile(c.ResponseWriter, c.Request, "hash-00/BenchmarkTest00073.html")
}

func (c *BenchmarkTest00073Controller) Post() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	theCookies := c.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, theCookie := range theCookies {
		if theCookie.Name == "BenchmarkTest00073" {
			param, _ = url.QueryUnescape(theCookie.Value)
			break
		}
	}

	var bar string
	guess := "ABC"
	switchTarget := guess[1]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	hash := md5.Sum([]byte(bar))
	hashString := hex.EncodeToString(hash[:])

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(c.ResponseWriter, "Unable to open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("hash_value=" + hashString + "\n"); err != nil {
		http.Error(c.ResponseWriter, "Unable to write to file", http.StatusInternalServerError)
		return
	}

	output, _ := json.Marshal(map[string]string{"message": "Sensitive value hashed and stored"})
	c.ResponseWriter.Write(output)
}
