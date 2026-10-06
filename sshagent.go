// The sshagent package provides an OS independent way to connect to a running "ssh-agent" process
// which returns an *[Agent] that is a wrapper for [agent.ExtendedAgent].
//
// On Windows, named pipes are used to connect to a local "ssh-agent", while on other platforms
// the "SSH_AUTH_SOCK" environment variable is expected to contain the path to a unix socket
// in order to communicate with the running "ssh-agent".
package sshagent

import (
	"net"

	"golang.org/x/crypto/ssh/agent"
)

// Agent wraps [agent.ExtendedAgent]
type Agent struct {
	agent.ExtendedAgent
	conn net.Conn
}

// NewAgent connects to a local SSH agent either via unix socket or named pipes depending on the OS
//
// In the case of a non-Windows OS, the "SSH_AUTH_SOCK" environment variable must be set or the
// process will fail.
//
// On Windows a connection is attempted to "\\.\pipe\openssh-ssh-agent" which is the default
// named pipe path for the native OpenSSH Authentication Agent.
func NewAgent() (*Agent, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}

	client := agent.NewClient(conn)

	return &Agent{client, conn}, nil
}

// Close closes the underlying connection to the SSH agent.
//
// The [Agent] must not be used after Close has been called.
func (a *Agent) Close() error {
	return a.conn.Close()
}
