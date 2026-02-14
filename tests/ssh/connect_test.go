package ssh_test

import (
	"bufio"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"shop.gyeongho.dev/internal/sshsrv"

	"github.com/charmbracelet/wish/testsession"
	gossh "golang.org/x/crypto/ssh"
)

// TestConnectAndTUIOutput checks that an SSH session shows TUI output containing the header (shop or terminal).
func TestConnectAndTUIOutput(t *testing.T) {
	keyPath := filepath.Join(t.TempDir(), "id_ed25519")
	srv, err := sshsrv.NewServer("127.0.0.1:0", keyPath, sshsrv.ShopHandler)
	if err != nil {
		t.Fatalf("NewServer: %v", err)
	}
	addr := testsession.Listen(t, srv)

	cfg := &gossh.ClientConfig{
		User: "test",
		Auth: []gossh.AuthMethod{
			gossh.Password(""),
		},
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
	}
	sess, err := testsession.NewClientSession(t, addr, cfg)
	if err != nil {
		t.Fatalf("NewClientSession: %v", err)
	}
	defer sess.Close()

	stdout, err := sess.StdoutPipe()
	if err != nil {
		t.Fatalf("StdoutPipe: %v", err)
	}
	if err := sess.RequestPty("xterm", 80, 24, nil); err != nil {
		t.Fatalf("RequestPty: %v", err)
	}
	if err := sess.Shell(); err != nil {
		t.Fatalf("Shell: %v", err)
	}
	rd := bufio.NewReader(stdout)
	var out strings.Builder
	for i := 0; i < 50; i++ {
		line, err := rd.ReadString('\n')
		if err != nil {
			break
		}
		out.WriteString(line)
		if strings.Contains(strings.ToLower(line), "shop") || strings.Contains(strings.ToLower(line), "terminal") {
			break
		}
	}
	got := out.String()
	if !strings.Contains(strings.ToLower(got), "shop") && !strings.Contains(strings.ToLower(got), "terminal") {
		t.Errorf("TUI output should contain 'shop' or 'terminal'; got (first 500 chars): %q", trunc(got, 500))
	}
}

// TestConnectAndFullFlowViewProductAndAddToCart checks that one SSH session can open Shop, add a product to cart, and open Cart.
func TestConnectAndFullFlowViewProductAndAddToCart(t *testing.T) {
	keyPath := filepath.Join(t.TempDir(), "id_ed25519")
	srv, err := sshsrv.NewServer("127.0.0.1:0", keyPath, sshsrv.ShopHandler)
	if err != nil {
		t.Fatalf("NewServer: %v", err)
	}
	addr := testsession.Listen(t, srv)

	cfg := &gossh.ClientConfig{
		User: "test",
		Auth: []gossh.AuthMethod{
			gossh.Password(""),
		},
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
	}
	sess, err := testsession.NewClientSession(t, addr, cfg)
	if err != nil {
		t.Fatalf("NewClientSession: %v", err)
	}
	defer sess.Close()

	stdout, err := sess.StdoutPipe()
	if err != nil {
		t.Fatalf("StdoutPipe: %v", err)
	}
	stdin, err := sess.StdinPipe()
	if err != nil {
		t.Fatalf("StdinPipe: %v", err)
	}
	if err := sess.RequestPty("xterm", 80, 24, nil); err != nil {
		t.Fatalf("RequestPty: %v", err)
	}
	if err := sess.Shell(); err != nil {
		t.Fatalf("Shell: %v", err)
	}
	rd := bufio.NewReader(stdout)
	var out strings.Builder
	for i := 0; i < 50; i++ {
		line, err := rd.ReadString('\n')
		if err != nil {
			break
		}
		out.WriteString(line)
		if strings.Contains(strings.ToLower(line), "shop") || strings.Contains(strings.ToLower(line), "terminal") {
			break
		}
	}
	time.Sleep(300 * time.Millisecond)
	if _, err := stdin.Write([]byte("a\rc")); err != nil {
		t.Fatalf("write keys: %v", err)
	}
	time.Sleep(400 * time.Millisecond)
	for i := 0; i < 250; i++ {
		line, err := rd.ReadString('\n')
		if err != nil {
			break
		}
		out.WriteString(line)
		if strings.Contains(line, "Total:") {
			break
		}
	}
	got := out.String()
	if !strings.Contains(got, "Oyster") {
		t.Errorf("full flow should show product (Oyster); len=%d, first 1500 chars: %q", len(got), trunc(got, 1500))
	}
	if !strings.Contains(got, "Total:") {
		t.Errorf("full flow should show Cart with Total: (add to cart); len=%d, first 1500 chars: %q", len(got), trunc(got, 1500))
	}
}

func trunc(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}
