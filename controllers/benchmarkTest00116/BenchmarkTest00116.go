package controllers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Employees struct {
	Employees []Employee `xml:"Employee"`
}

type Employee struct {
	ID   string `xml:"emplid,attr"`
	Data string `xml:",chardata"`
}

func BenchmarkTest00116(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest00116",
			Value:  "2222",
			Path:   r.RequestURI,
			Domain: getDomain(r),
			Secure: true,
			MaxAge: 60 * 3,
		})
		http.ServeFile(w, r, "xpathi-00/BenchmarkTest00116.html")
		return
	}

	if r.Method == http.MethodPost {
		cookies := r.Cookies()
		param := "noCookieValueSupplied"
		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest00116" {
				param = cookie.Value
				break
			}
		}

		bar := "safe!"
		map51005 := make(map[string]interface{})
		map51005["keyA-51005"] = "a_Value"
		map51005["keyB-51005"] = param
		map51005["keyC"] = "another_Value"

		bar = map51005["keyB-51005"].(string)
		bar = map51005["keyA-51005"].(string)

		file, err := os.Open("employees.xml")
		if err != nil {
			http.Error(w, "Error opening XML file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		data, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, "Error reading XML file", http.StatusInternalServerError)
			return
		}

		var employees Employees
		if err := xml.Unmarshal(data, &employees); err != nil {
			http.Error(w, "Error parsing XML", http.StatusInternalServerError)
			return
		}

		for _, employee := range employees.Employees {
			if employee.ID == bar {
				fmt.Fprintf(w, "Your query results are: <br/>%s<br/>", employee.Data)
			}
		}
	}
}

func getDomain(r *http.Request) string {
	return r.URL.Host
}

func main() {
	http.HandleFunc("/xpathi-00/BenchmarkTest00116", BenchmarkTest00116)
	http.ListenAndServe(":8080", nil)
}
