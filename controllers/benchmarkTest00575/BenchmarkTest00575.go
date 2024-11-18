package controllers

import (
	"net/http"
	"os/exec"
	"strings"
)

type BenchmarkTest00575 struct{}

func (b *BenchmarkTest00575) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := ""

	for name, values := range r.Form {
		for _, value := range values {
			if value == "BenchmarkTest00575" {
				param = name
				break
			}
		}
	}

	bar := ""
	if param != "" {
		valuesList := []string{"safe"}
		valuesList = append(valuesList, param)
		valuesList = append(valuesList, "moresafe")

		valuesList = valuesList[1:] // remove the 1st safe value

		bar = valuesList[0] // get the param value
	}

	cmd := ""
	if strings.Contains(strings.ToLower(r.UserAgent()), "windows") {
		cmd = "echo"
	}

	if cmd != "" && bar != "" {
		out, err := exec.Command(cmd, bar).CombinedOutput()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(out)
	}
}

func main() {
	http.Handle("/cmdi-00/BenchmarkTest00575", &BenchmarkTest00575{})
	http.ListenAndServe(":8080", nil)
}
