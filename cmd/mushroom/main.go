package main

import (
	"context"
	"fmt"
	"os"

	"mushroom.gyeongho.dev/internal/apiclient"
	"mushroom.gyeongho.dev/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	client := apiclient.MockClient{}
	m := tui.NewModel(client)
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
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
