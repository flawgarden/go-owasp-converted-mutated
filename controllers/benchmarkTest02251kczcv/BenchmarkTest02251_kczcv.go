package controllers

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

type BenchmarkTest02251 struct{}

func (b *BenchmarkTest02251) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var param string
	if values, ok := r.Form["BenchmarkTest02251"]; ok && len(values) > 0 {
		param = values[0]
	}

	bar := b.doSomething(param)

	var args []string

str123 := "str123"
addSuffix(&str123, "_suffix")
bar = str123

	osName := getOSName()

	if strings.Contains(osName, "Windows") {
		args = []string{"cmd.exe", "/c", "echo " + bar}
	} else {
		args = []string{"sh", "-c", "ping -c1 " + bar}
	}

	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, fmt.Sprintf("Problem executing command: %s", err), http.StatusInternalServerError)
		return
	}
	w.Write(out)
}

func (b *BenchmarkTest02251) doSomething(param string) string {
	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}

func getOSName() string {
	return "Linux" // or use runtime.GOOS for actual OS detection in Go
}

func addSuffix(s *string, suf string) {
	*s = *s + suf
}

func addSuffixDoublePointer(s **string, suf *string) {
	**s = **s + *suf
}

func addSuffixDoublePointerBroken(s **string, suf *string) {
	*s = new(string)
	**s = **s + *suf
}

func swapStrings(a, b *string) {
	temp := *a
	*a = *b
	*b = temp
}

func removeSpaces(s *string) {
    *s = strings.ReplaceAll(*s, " ", "")
}


