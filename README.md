# sshagent

[![Go Report Card](https://goreportcard.com/badge/github.com/andrewheberle/sshagent?logo=go&style=flat-square)](https://goreportcard.com/report/github.com/andrewheberle/sshagent)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/andrewheberle/sshagent)
[![codecov](https://codecov.io/gh/andrewheberle/sshagent/graph/badge.svg?token=MNFPOWU3VV)](https://codecov.io/gh/andrewheberle/sshagent)

This package provides an OS independent way to connect to a running "ssh-agent" process
which returns an `sshagent.*Agent` that is a wrapper for [golang.org/x/crypto/ssh/agent.ExtendedAgent](https://pkg.go.dev/golang.org/x/crypto/ssh/agent#ExtendedAgent).

On Windows, named pipes are used to connect to a local "ssh-agent", while on other platforms
the "SSH_AUTH_SOCK" environment variable is expected to contain the path to a unix socket
in order to communicate with the running "ssh-agent".

## Example

```go
package main

import (
	"fmt"
    "os"

	"github.com/andrewheberle/sshagent"
)

func main() {
	client, err := sshagent.NewAgent()
	if err != nil {
		fmt.Printf("error connecting to agent")
        os.Exit(1)
	}

	if _, err := client.List(); err != nil {
		fmt.Printf("error listing keys from agent")
	}
}
```
