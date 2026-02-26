package labels

import (
	"github.com/adm87/msgen/styles/palette"
	"github.com/charmbracelet/lipgloss"
)

var ErrorLabel = lipgloss.NewStyle().
	Foreground(palette.Red).
	Bold(true)

func Error(err error) string {
	return ErrorLabel.Render("Error: " + err.Error())
}
