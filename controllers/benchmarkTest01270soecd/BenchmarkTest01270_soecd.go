package controllers

import (
	"net/http"
	"os/exec"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01270 struct{}

func (b *BenchmarkTest01270) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01270")
	if param == "" {
		param = ""
	}

var a12341 BinaryOpInterfaceDefault2 = &BinaryOpInterfaceDefault2Implementation{}
param = a12341.InterfaceCall(param, "")

	bar := b.doSomething(param)

	var a1, a2 string
	if strings.Contains(strings.ToLower(r.UserAgent()), "windows") {
		a1 = "cmd.exe"
		a2 = "/c"
	} else {
		a1 = "sh"
		a2 = "-c"
	}
	args := []string{a1, a2, "echo " + bar}

	process := exec.Command(args[0], args[1:]...)
	output, err := process.CombinedOutput()
	if err != nil {
		http.Error(w, "Problem executing command", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write(output)
}

func (b *BenchmarkTest01270) doSomething(param string) string {
	bar := ""
	num := 106

	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}
	return bar
}
