package tasks

import (
	"os/exec"
	"strings"
	"time"

	"github.com/adm87/msgen/app"
	tea "github.com/charmbracelet/bubbletea"
)

type InitModuleStartMsg struct{}
type InitModuleCompletedMsg struct {
	Output string
	Err    error
}

type InitModuleTask struct {
	status app.TaskStatus

	path   string
	module string
	view   string
}

func InitModule(path, module string) *InitModuleTask {
	return &InitModuleTask{
		path:   path,
		module: module,
	}
}

func (t *InitModuleTask) Title() string {
	return "Running 'go mod init'"
}

func (t *InitModuleTask) Status() app.TaskStatus {
	return t.status
}

func (t *InitModuleTask) Init() tea.Cmd {
	t.status = app.TaskStatusInProgress
	return func() tea.Msg { return InitModuleStartMsg{} }
}

func (t *InitModuleTask) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case InitModuleStartMsg:
		return tea.Tick(app.TaskWaitBuffer, func(_ time.Time) tea.Msg {
			cmd := exec.Command("go", "mod", "init", t.module)
			cmd.Dir = t.path

			output, err := cmd.CombinedOutput()
			return InitModuleCompletedMsg{
				Output: strings.TrimSpace(string(output)),
				Err:    err,
			}
		})

	case InitModuleCompletedMsg:
		return tea.Tick(app.TaskWaitBuffer, func(_ time.Time) tea.Msg {
			if msg.Err != nil {
				t.view = "failed to run 'go mod init'"
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

func (t *InitModuleTask) View() string {
	if t.view != "" {
		return t.view + "\n"
	}
	return ""
}
