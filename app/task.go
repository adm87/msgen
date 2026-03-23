package app

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	TaskWaitBuffer = 250 * time.Millisecond
)

type TaskCompletedMsg struct {
	Task Task
	Err  error
}

func TaskCompleted(task Task, err error) tea.Cmd {
	return func() tea.Msg {
		return TaskCompletedMsg{
			Task: task,
			Err:  err,
		}
	}
}

type Task interface {
	Title() string
	Status() TaskStatus

	Init() tea.Cmd
	Update(msg tea.Msg) tea.Cmd
	View() string
}
