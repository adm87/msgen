package app

import (
	"context"
	"strings"

	"github.com/adm87/msgen/palette"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	spinner spinner.Model

	tasks       []Task
	currentTask int
	buffer      *strings.Builder
	err         error
}

func Run(ctx context.Context, tasks []Task) error {
	m := &model{
		tasks:   tasks,
		buffer:  &strings.Builder{},
		spinner: palette.NewProgressSpinner(),
	}

	if _, err := tea.NewProgram(m, tea.WithContext(ctx)).Run(); err != nil {
		return err
	}

	return m.err
}

func (m *model) Init() tea.Cmd {
	if len(m.tasks) == 0 {
		return tea.Quit
	}
	return tea.Batch(m.tasks[0].Init(), m.spinner.Tick)
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case TaskCompletedMsg:
		if msg.Err != nil {
			m.err = msg.Err
			return m, tea.Quit
		}

		m.currentTask++

		if m.currentTask >= len(m.tasks) {
			return m, tea.Quit
		}

		cmd := m.tasks[m.currentTask].Init()
		return m, cmd
	}

	if m.currentTask < len(m.tasks) {
		cmd := m.tasks[m.currentTask].Update(msg)
		return m, cmd
	}

	return m, tea.Quit
}

func (m *model) View() string {
	m.buffer.Reset()

	for i, task := range m.tasks {
		if i > m.currentTask {
			break // Don't render tasks that haven't started yet
		}

		icon := GetStatusIcon(task.Status())
		if task.Status() == TaskStatusInProgress {
			icon = strings.TrimSpace(m.spinner.View())
		}

		title := task.Title()
		view := task.View()

		if i == m.currentTask {
			m.buffer.WriteString(icon + " " + palette.Render(palette.TitleActiveStyle, title) + "\n")
		} else {
			m.buffer.WriteString(icon + " " + palette.Render(palette.TitlePendingStyle, title) + "\n")
		}

		if view != "" {
			m.buffer.WriteString(view + "\n")
		}
	}

	return m.buffer.String()
}
