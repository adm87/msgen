package tasks

import (
	"github.com/adm87/msgen/app"
	tea "github.com/charmbracelet/bubbletea"
)

type GenerateSpecTask struct {
	status app.TaskStatus
}

func GenerateSpec(path string) *GenerateSpecTask {
	return &GenerateSpecTask{
		status: app.TaskStatusPending,
	}
}

func (t *GenerateSpecTask) Title() string {
	return "Generating service specification"
}

func (t *GenerateSpecTask) Status() app.TaskStatus {
	return t.status
}

func (t *GenerateSpecTask) Init() tea.Cmd {
	return func() tea.Msg {
		t.status = app.TaskStatusInProgress
		return nil
	}
}

func (t *GenerateSpecTask) Update(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

func (t *GenerateSpecTask) View() string {
	return ""
}
