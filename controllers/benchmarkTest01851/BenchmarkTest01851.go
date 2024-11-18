package controllers

import (
	"net/http"
	"os/exec"
	"strings"
)

type BenchmarkTest01851 struct{}

func (b *BenchmarkTest01851) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	http.SetCookie(w, &http.Cookie{
		Name:   "BenchmarkTest01851",
		Value:  "ECHOOO",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   r.URL.Path,
		Domain: r.Host,
	})
	http.ServeFile(w, r, "./cmdi-02/BenchmarkTest01851.html")
}

func (b *BenchmarkTest01851) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookie, err := r.Cookie("BenchmarkTest01851")
	param := "noCookieValueSupplied"
	if err == nil {
		param = cookie.Value
	}

	bar := doSomething(param)

	var argList []string
	if strings.Contains(strings.ToLower(r.UserAgent()), "windows") {
		argList = []string{"cmd.exe", "/c", "echo " + bar}
	} else {
		argList = []string{"sh", "-c", "echo " + bar}
	}

	cmd := exec.Command(argList[0], argList[1:]...)
	output, err := cmd.Output()
	if err != nil {
		http.Error(w, "Error executing command", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func doSomething(param string) string {
	bar := ""
	switchTarget := "C" // This simulates the switch case with a predetermined target
	switch switchTarget {
	case "A":
		bar = param
	case "B":
		bar = "bobs_your_uncle"
	case "C", "D":
		bar = param
	default:
		bar = "bobs_your_uncle"
	}
	return bar
}
