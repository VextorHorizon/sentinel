# Sentinel

A lightweight TCP network scanner written in Go.

Sentinel is a learning-focused project built to explore low-level networking, CLI development, and eventually concurrency in Go.

> **Current version:** `v0.0.3`  
> Sentinel is still in early development.

## Current Features

- Scan a TCP port on an IPv4 address
- Accept a target IP and port from command-line arguments
- Detect whether a TCP connection can be established
- Scan multiple targets sequentially
- Basic input validation for CLI arguments

## Usage

Build Sentinel:

```bash
go build -o sentinel
```

Run it with a target IP and port:

```bash
./sentinel <ip> <port>
```

Example:

```bash
./sentinel 127.0.0.1 8080
```

On Windows:

```powershell
.\sentinel.exe 127.0.0.1 8080
```

Example output:

```text
127.0.0.1:8080
Port 8080 is open!
```

or:

```text
127.0.0.1:8080
Port 8080 is closed!
```

## Current Architecture

Targets are represented using an `IPAddress` struct:

```go
type IPAddress struct {
    ip   string
    port int
}
```

Each target is passed into:

```go
Scanner(target IPAddress)
```

The scanner builds an `IP:PORT` address and attempts a TCP connection using Go's `net` package.

At the moment, targets are scanned **sequentially**, meaning a slow connection attempt can delay every target after it.

The current development build also contains several hard-coded targets used for testing alongside the target supplied through the CLI.

## Roadmap

Planned improvements include:

- Remove hard-coded development targets
- Scan multiple ports
- Scan multiple hosts
- Add configurable connection timeout
- Add concurrent scanning with goroutines
- Improve error handling and input validation
- Clean up connection lifecycle
- Improve CLI output
- Split the project into multiple files/packages as it grows

## Why Sentinel?

Sentinel is primarily a project for learning Go by building something real instead of only following isolated exercises.

The project is being developed incrementally, with each version introducing new Go and networking concepts while keeping the implementation understandable.

Long-term, the goal is to evolve Sentinel from a simple TCP connection checker into a small, practical network scanning CLI.

## Status

Early development / learning project.

Expect breaking changes as the architecture evolves.
