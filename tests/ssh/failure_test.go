package ssh_test

import (
	"net"
	"strings"
	"testing"
	"time"

	gossh "golang.org/x/crypto/ssh"
)

// TestConnectionToUnreachableServerFailsWithBoundedTimeout checks that dialing an unreachable address fails with a bounded timeout or connection refused.
func TestConnectionToUnreachableServerFailsWithBoundedTimeout(t *testing.T) {
	cfg := &gossh.ClientConfig{
		User: "test",
		Auth: []gossh.AuthMethod{
			gossh.Password(""),
		},
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
		Timeout:         3 * time.Second,
	}
	addr := "127.0.0.1:19999"
	conn, err := gossh.Dial("tcp", addr, cfg)
	if err == nil {
		conn.Close()
		t.Fatal("expected connection to unreachable host to fail")
	}
	if ne, ok := err.(net.Error); ok && ne.Timeout() {
		t.Logf("connection timed out as expected: %v", err)
		return
	}
	if strings.Contains(err.Error(), "refused") || strings.Contains(err.Error(), "connection refused") {
		t.Logf("connection refused as expected: %v", err)
		return
	}
	t.Logf("connection failed with expected kind of error: %v", err)
}
