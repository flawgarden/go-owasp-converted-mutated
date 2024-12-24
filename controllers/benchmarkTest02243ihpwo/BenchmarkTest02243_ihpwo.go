package controllers

import (
	"net/http"
	"os/exec"
)

type BenchmarkTest02243 struct{}

func (b *BenchmarkTest02243) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02243")

	bar := doSomething(r, param)

tmpArrayUnique42 := []string{"", "", "", ""}
tmpArrayUnique42[0] = bar
ah := NewArrayHolderWithValues(tmpArrayUnique42)
bar = ah.Values[0]

	var argList []string
	if isWindows() {
		argList = []string{"cmd.exe", "/c", "echo " + bar}
	} else {
		argList = []string{"sh", "-c", "echo " + bar}
	}

	cmd := exec.Command(argList[0], argList[1:]...)
	output, err := cmd.Output()
	if err != nil {
		http.Error(w, "Problem executing command", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write(output)
}

func doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map19941 := make(map[string]interface{})
	map19941["keyA-19941"] = "a-Value"
	map19941["keyB-19941"] = param
	map19941["keyC"] = "another-Value"
	bar = map19941["keyB-19941"].(string)

	return bar
}

func isWindows() bool {
	return exec.Command("cmd", "/c", "ver").Run() == nil
}
