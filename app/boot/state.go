package boot

import (
	"github.com/adm87/msgen/data"
	"github.com/adm87/msgen/state"
	tea "github.com/charmbracelet/bubbletea"
)

const State state.State = "bootstrap"

func Update(ctx *data.MSGenContext, msg tea.Msg) tea.Cmd {
	switch msg.(type) {
	case state.EnterMsg:
		return state.SendMsg(state.NewUpdateMsg(State))
	}
	return state.SendMsg(state.NewCompletionMsg(State))
}

func View(ctx *data.MSGenContext) string {
	return "Bootstrapping... "
}
