package controllers

import (
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/cmdi-02/BenchmarkTest02496", BenchmarkTest02496)
	http.ListenAndServe(":8080", nil)
}

func BenchmarkTest02496(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	r.ParseForm()
	values := r.Form["BenchmarkTest02496"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := doSomething(r, param)

	var cmd string
	var args []string
	if cmd = "echo"; isWindows() {
		args = []string{"cmd.exe", "/c", bar}
	} else {
		args = []string{"sh", "-c", bar}
	}

	output, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		http.Error(w, "Problem executing cmdi - "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

func doSomething(r *http.Request, param string) string {
	bar := ""
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}

func isWindows() bool {
	return len(os.Getenv("OS")) > 0 && os.Getenv("OS") == "Windows_NT"
}
