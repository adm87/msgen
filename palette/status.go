package palette

import "github.com/charmbracelet/lipgloss"

var (
	PendingStatus    = lipgloss.NewStyle().Foreground(lipgloss.Color("244")) // Grey
	InProgressStatus = lipgloss.NewStyle().Foreground(lipgloss.Color("214")) // Orange/Yellow
	SuccessStatus    = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))  // Green
	ErrorStatus      = lipgloss.NewStyle().Foreground(lipgloss.Color("196")) // Red
)

var (
	TitleActiveStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("33"))  // Blue
	TitlePendingStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("256")) // Default terminal color
)

func Render(style lipgloss.Style, text string) string {
	return style.Render(text)
}
