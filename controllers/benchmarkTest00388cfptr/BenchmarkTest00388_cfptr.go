package controllers

import (
"fmt"
"net/http"
"strings"
)

func BenchmarkTest00388(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest00388")

counter := func() func(str string) string {
    count := 0
    return func(str string) string {
        count++
        if count == 1 {
            return str
        } else {
            return "fixed_string"
        }

    }
}()
param = counter(param)

	if param == "" {
		param = ""
	}

	sbxyz30382 := strings.Builder{}
	sbxyz30382.WriteString(param)
	bar := sbxyz30382.String() + "_SafeStuff"

	w.Header().Set("X-XSS-Protection", "0")
	fmt.Fprintln(w, []rune(bar))
}

func main() {
	http.HandleFunc("/xss-00/BenchmarkTest00388", BenchmarkTest00388)
	http.ListenAndServe(":8080", nil)
}

func makePrefixer(prefix string) func(string) string {
    return func(value string) string {
        return fmt.Sprintf("%s%s", prefix, value)
    }
}

func makeMessageGenerator(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return prefix + name
		}
	}
}

func makeMessageGeneratorBroken(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return name
		}
	}
}


