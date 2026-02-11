package tui

import (
	"math/rand"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	loaderBoxWidth    = 40
	loaderBoxHeight   = 20
	loaderInnerWidth  = 40
	loaderInnerHeight = 5
	loaderInnerText   = "Loading mushroom.gyeongho.dev"
	loaderInnerPad    = 6
)

var (
	loaderOuterStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	loaderInnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true)
)

const loaderChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// Loader returns a loading view: a rectangle of random characters with a centered inner line of text.
func Loader() string {
	innerStartCol := (loaderBoxWidth - loaderInnerWidth) / 2
	innerStartRow := (loaderBoxHeight - loaderInnerHeight) / 2
	innerEndCol := innerStartCol + loaderInnerWidth
	innerEndRow := innerStartRow + loaderInnerHeight
	middleRow := innerStartRow + loaderInnerHeight/2

	textLen := len(loaderInnerText)
	padLeft := (loaderInnerWidth - textLen) / 2
	if padLeft > loaderInnerPad {
		padLeft = loaderInnerPad
	}
	padRight := loaderInnerWidth - padLeft - textLen
	if padRight < 0 {
		padRight = 0
	}
	middleLineContent := strings.Repeat(" ", padLeft) + loaderInnerText + strings.Repeat(" ", padRight)
	if len(middleLineContent) > loaderInnerWidth {
		middleLineContent = middleLineContent[:loaderInnerWidth]
	}

	var b strings.Builder
	for row := 0; row < loaderBoxHeight; row++ {
		for col := 0; col < loaderBoxWidth; col++ {
			inInner := col >= innerStartCol && col < innerEndCol && row >= innerStartRow && row < innerEndRow
			if inInner && row == middleRow {
				idx := col - innerStartCol
				if idx < len(middleLineContent) {
					b.WriteByte(middleLineContent[idx])
				} else {
					b.WriteByte(' ')
				}
			} else if inInner {
				b.WriteByte(' ')
			} else {
				b.WriteByte(loaderChars[rand.Intn(len(loaderChars))])
			}
		}
		if row < loaderBoxHeight-1 {
			b.WriteByte('\n')
		}
	}
	raw := b.String()
	lines := strings.Split(raw, "\n")
	for i := range lines {
		if i >= innerStartRow && i < innerEndRow {
			if i == middleRow {
				lines[i] = loaderInnerStyle.Render(lines[i])
			} else {
				lines[i] = loaderOuterStyle.Render(lines[i])
			}
		} else {
			lines[i] = loaderOuterStyle.Render(lines[i])
		}
	}
	block := strings.Join(lines, "\n")
	const verticalPadLines = 5
	pad := strings.Repeat("\n", verticalPadLines)
	return pad + block + pad
}
