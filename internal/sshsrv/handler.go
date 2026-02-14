package sshsrv

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"

	"shop.gyeongho.dev/internal/apiclient"
	"shop.gyeongho.dev/internal/tui"
)

// ShopHandler returns a full shop TUI model for the SSH session with products, about, and FAQ pre-loaded.
func ShopHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, _ := s.Pty()
	width := pty.Window.Width
	height := pty.Window.Height
	if width <= 0 {
		width = 80
	}
	if height <= 0 {
		height = 24
	}

	client := apiclient.MockClient{}
	m := tui.NewModel(client)
	m.Width = width
	m.Height = height

	ctx := context.Background()
	if prods, err := client.GetProducts(ctx); err == nil {
		m.Products = prods
	} else {
		m.Error = err.Error()
	}
	if about, err := client.GetAbout(ctx); err == nil {
		m.About = about
	} else if m.Error == "" {
		m.Error = err.Error()
	}
	if faq, err := client.GetFAQ(ctx); err == nil {
		m.FAQ = faq
	} else if m.Error == "" {
		m.Error = err.Error()
	}
	m.Loading = false

	return m, []tea.ProgramOption{tea.WithAltScreen()}
}
