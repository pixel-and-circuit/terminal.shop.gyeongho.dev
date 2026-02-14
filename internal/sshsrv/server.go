package sshsrv

import (
	"time"

	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
)

const (
	defaultIdleTimeout = 30 * time.Second
	defaultMaxTimeout  = 30 * time.Second
)

// NewServer creates a Wish SSH server that serves the given Bubble Tea handler.
// addr is e.g. ":2222" or "127.0.0.1:0" for tests; hostKeyPath is e.g. ".ssh/id_ed25519" (Wish creates if missing).
// handler must be non-nil.
func NewServer(addr, hostKeyPath string, handler bubbletea.Handler) (*ssh.Server, error) {
	return wish.NewServer(
		wish.WithAddress(addr),
		wish.WithHostKeyPath(hostKeyPath),
		wish.WithIdleTimeout(defaultIdleTimeout),
		wish.WithMaxTimeout(defaultMaxTimeout),
		ssh.PublicKeyAuth(func(ssh.Context, ssh.PublicKey) bool { return true }),
		ssh.PasswordAuth(func(ssh.Context, string) bool { return true }),
		wish.WithMiddleware(
			bubbletea.Middleware(handler),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
}
