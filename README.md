# c2go

**c2go** is a minimal, cross-platform command-and-control (C2) server and matching implant, built in Go. It's designed to be as small and simple as possible — perfect for single-agent R&D, purple team payload testing, or detection tuning on a local network.

> Leave the victim machine running in another room and test from anywhere on the LAN — no need to bring both laptops. 

This project was built to make local agent testing dead simple, portable, and stealthy — without needing a full offensive infrastructure setup.

---

## ✨ Features

- 🖥️ Simple terminal-based C2 server with a clean `cmd>` interface
- 🔁 HTTP beaconing agent (cross-compiled for macOS and Linux)
- 📡 Custom server IP can be passed at build time via `-ldflags`
- 🔐 Minimal footprint: single binary, no dependencies
- ⚙️ Built for purple team testing, payload R&D, and EDR detection validation

---

## 🚀 Quick Start

### 1. Clone and initialize the module

```bash
git clone https://github.com/HackerIndustrial-Tooling/C2go.git
cd c2go
go mod init c2go
```

### 2. Run the C2 server

```bash
go run server.go
```

This starts the server on `http://0.0.0.0:8080` and drops you into an interactive prompt (`cmd>`). When the agent connects, it will display connection logs and any command results.

---

## 🛠️ Build the Agent

You can pass the C2 server IP/port at **build time** using Go's `-ldflags`.

### 🔧 Build for macOS (Intel):

```bash
GOOS=darwin GOARCH=amd64 go build -o agent-macos-x86 \
  -ldflags="-X 'main.serverURL=http://192.168.1.100:8080'" \
  agent/main.go
```

### 🍎 Build for macOS (Apple Silicon):

```bash
GOOS=darwin GOARCH=arm64 go build -o agent-macos-arm64 \
  -ldflags="-X 'main.serverURL=http://192.168.1.100:8080'" \
  agent/main.go
```

> Replace `192.168.1.100` with the IP of the machine running the server.

---

## 🧪 Usage

- Launch the C2 server (`go run server.go`)
- Run the compiled agent on another machine on the same LAN
- Enter shell commands at the `cmd>` prompt
- See results stream back in real time

---

## 🔧 Agent Notes

- The agent polls `/get` every few seconds
- When it receives a command, it runs it via `bash -c` and POSTs the result to `/result`
- Server logs the first connection from any unique IP

---

## 🧠 Why?

This is the perfect local dev tool for:

- R&D with new payloads
- Rapid EDR testing in isolated labs
- Purple team operator tools
- Fast iteration without needing a full C2 framework

---

## 📜 License

MIT

---
