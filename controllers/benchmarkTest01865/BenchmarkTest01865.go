package controllers

import (
	"net/http"
	"net/url"
	"os/exec"
	"strings"
)

type BenchmarkTest01865 struct{}

func (b *BenchmarkTest01865) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{Name: "BenchmarkTest01865", Value: "ls", Path: r.URL.Path, MaxAge: 60 * 3, Secure: true})
		http.ServeFile(w, r, "cmdi-02/BenchmarkTest01865.html")
		return
	}

	if r.Method == http.MethodPost {
		cookies := r.Cookies()
		param := "noCookieValueSupplied"

		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest01865" {
				param, _ = url.QueryUnescape(cookie.Value)
				break
			}
		}

		bar := doSomething(param)
		var cmd string
		if strings.Contains(strings.ToLower(r.UserAgent()), "windows") {
			cmd = "cmd /C echo "
		} else {
			cmd = "echo "
		}

		output, err := exec.Command("sh", "-c", cmd+bar).Output()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(output)
	}
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}
	return bar
}
