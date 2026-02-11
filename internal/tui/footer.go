package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const footerWidth = 60

var (
	footerBarStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	footerTextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

// RenderFooter returns the footer: a horizontal line and centered hint text.
func RenderFooter() string {
	bar := strings.Repeat("─", footerWidth)
	text := "+/- qty   c cart   q quit"
	textLine := lipgloss.PlaceHorizontal(footerWidth, lipgloss.Center, footerTextStyle.Render(text))
	return footerBarStyle.Render(bar) + "\n" + textLine
}
