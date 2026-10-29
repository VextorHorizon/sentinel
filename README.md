# Sentinel

A lightweight TCP network scanner written in Go.

Sentinel is a learning-focused project built to explore low-level networking, CLI development, and eventually concurrency in Go.

> **Current version:** `v0.1.0`  
> Sentinel is still in development as learning project.

## Current Features

- Scan TCP ports on an IPv4 address
- Accept a target IP and port from command-line arguments
- Detect whether a TCP connection can be established
- Basic input validation for CLI arguments

## Usage

Build Sentinel:

```bash
go build -o sentinel
```

Run it with a target IP and port:

```bash
./sentinel <option> <port> <ip>
```

Example:

```bash
./sentinel -p 5350 127.0.0.1
./sentinel -p 5350-5400 127.0.0.1
./sentinel -p 5350,5351,5352 127.0.0.1
```

On Windows:

```powershell
.\sentinel.exe -p 5350 127.0.0.1
.\sentinel.exe -p 5350,5351,5352 127.0.0.1
```

Example output:

```text
5354             open
5350             closed
5351             closed
```

## Current Architecture

Targets are represented using an `IPAddress` struct:

```go
type IPAddress struct {
    ip   string
    port []int
}
```

Each target is passed into:

```go
Scanner(target IPAddress)
```

The scanner builds an `IP:PORT` address and attempts a TCP connection using Go's `net` package.

## Roadmap

Planned improvements include:

- Remove hard-coded development targets [x]
- Scan multiple ports [x]
- Parser range input [x]
- Scan multiple hosts [] (v0.2)
- Add configurable connection timeout[x]
- Add concurrent scanning with goroutines [x]
- Improve error handling and input validation 
- Clean up connection lifecycle [x]
- Improve CLI output
- Split the project into multiple files/packages as it grows

## Why Sentinel?

Sentinel is primarily a project for learning Go by building something real instead of only following isolated exercises.

The project is being developed incrementally, with each version introducing new Go and networking concepts while keeping the implementation understandable.

Long-term, the goal is to evolve Sentinel from a simple TCP connection checker into a small, practical network scanning CLI.

## Status

Stable v0.1.0 version and still in development / learning project.

Expect breaking changes as the architecture evolves.
