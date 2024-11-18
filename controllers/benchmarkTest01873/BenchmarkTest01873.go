package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest01873 struct{}

func (bt *BenchmarkTest01873) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest01873",
			Value:  "my_user_id",
			Path:   r.URL.Path,
			Secure: true,
		})
		http.ServeFile(w, r, "trustbound-01/BenchmarkTest01873.html")
	} else if r.Method == http.MethodPost {
		cookies := r.Cookies()
		param := "noCookieValueSupplied"
		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest01873" {
				decodedValue, _ := url.QueryUnescape(cookie.Value)
				param = decodedValue
				break
			}
		}

		bar := doSomething(param)
		r.AddCookie(&http.Cookie{Name: "userid", Value: bar})

		fmt.Fprintf(w, "Item: 'userid' with value: '%s' saved in session.", bar)
	}
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = append(valuesList[:0], valuesList[1:]...) // remove the 1st safe value
		bar = valuesList[0]                                    // get the last 'safe' value
	}
	return bar
}

func main() {
	http.Handle("/trustbound-01/BenchmarkTest01873", &BenchmarkTest01873{})
	http.ListenAndServe(":8080", nil)
}
