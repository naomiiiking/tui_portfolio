# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this project is

A personal SSH terminal website — people connect with `ssh <hostname>` and get an interactive TUI showing about, experience, and links. Built with Go + the Charmbracelet stack.

## Commands

```bash
# Run locally (server starts on port 2222)
go run .

# Connect to local server (in a second terminal)
ssh localhost -p 2222

# Build binary
go build -o terminal-website .

# Add/update dependencies
go get <package>@latest
go mod tidy
```

Go must be on your PATH (`/opt/homebrew/bin/go` on this machine).

## Architecture

```
main.go              — SSH server setup (Wish); maps each SSH session → Bubbletea program
ui/model.go          — Bubbletea Model: state, keyboard handling, and all view rendering
ui/styles.go         — Lip Gloss style variables (colors, bold, underline)
content/content.go   — All user-facing content: name, bio, jobs, links
```

**Data flow:** `main.go` starts a Wish SSH server → each connection spawns a `tea.Program` via `teaHandler` → Bubbletea runs the `ui.Model` update/view loop → Lip Gloss renders styled strings to the SSH PTY.

**The only file that needs editing for content changes is `content/content.go`.**

## Key libraries

| Library | Role |
|---|---|
| `charmbracelet/wish` | SSH server that hands connections to a Bubbletea program |
| `charmbracelet/bubbletea` | Elm-style TUI framework: `Init / Update / View` |
| `charmbracelet/lipgloss` | Terminal styling — colours, padding, width/alignment |

## SSH host key

On first run, `wish` auto-generates a host key at `.ssh/id_ed25519` (gitignored). Don't commit this file.

## Deployment

The binary needs to run on a public server with port 2222 (or 22) open. Typical setup:
- Build for Linux: `GOOS=linux GOARCH=amd64 go build -o terminal-website .`
- Copy binary + run as a systemd service or with a process manager
- Open the SSH port in your firewall/security group
