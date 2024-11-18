package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00494Controller struct {
	web.Controller
}

func (c *BenchmarkTest00494Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00494Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00494")
	bar := "alsosafe"

	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}

	var cmd string
	var args []string
	osName := getOSName()

	if strings.Contains(osName, "Windows") {
		args = []string{"cmd.exe", "/c", cmd + bar}
	} else {
		args = []string{"sh", "-c", cmd + bar}
	}

	executeCommand(args, c.Ctx.ResponseWriter)
}

func getOSName() string {
	return "Some OS Name" // Replace with actual OS detection if needed
}

func executeCommand(args []string, w http.ResponseWriter) {
	cmd := strings.Join(args, " ")
	// This is where the command would be executed
	json.NewEncoder(w).Encode(map[string]string{"command": cmd}) // Just an example
}
