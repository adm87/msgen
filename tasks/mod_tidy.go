package tasks

import (
	"os/exec"
	"strings"
	"time"

	"github.com/adm87/msgen/app"
	tea "github.com/charmbracelet/bubbletea"
)

type ModTidyStartMsg struct{}
type ModTidyCompletedMsg struct {
	Output string
	Err    error
}

type ModTidyTask struct {
	status app.TaskStatus

	path string
	view string
}

func ModTidy(path string) *ModTidyTask {
	return &ModTidyTask{
		path: path,
	}
}

func (t *ModTidyTask) Title() string {
	return "Running 'go mod tidy'"
}

func (t *ModTidyTask) Status() app.TaskStatus {
	return t.status
}

func (t *ModTidyTask) Init() tea.Cmd {
	t.status = app.TaskStatusInProgress
	return func() tea.Msg { return ModTidyStartMsg{} }
}

func (t *ModTidyTask) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case ModTidyStartMsg:
		return tea.Tick(app.TaskWaitBuffer, func(_ time.Time) tea.Msg {
			cmd := exec.Command("go", "mod", "tidy")
			cmd.Dir = t.path

			output, err := cmd.CombinedOutput()
			return ModTidyCompletedMsg{
				Output: strings.TrimSpace(string(output)),
				Err:    err,
			}
		})

	case ModTidyCompletedMsg:
		return tea.Tick(app.TaskWaitBuffer, func(_ time.Time) tea.Msg {
			if msg.Err != nil {
				t.view = "failed to run 'go mod tidy'"
				t.status = app.TaskStatusFailed
			} else {
				t.view = msg.Output
				t.status = app.TaskStatusCompleted
			}
			return app.TaskCompletedMsg{
				Task: t,
				Err:  msg.Err,
			}
		})
	}
	return nil
}

func (t *ModTidyTask) View() string {
	if t.view != "" {
		return t.view + "\n"
	}
	return ""
}
