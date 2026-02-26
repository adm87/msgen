package app

import (
	"fmt"
	"strings"

	"github.com/adm87/msgen/app/boot"
	"github.com/adm87/msgen/data"
	"github.com/adm87/msgen/state"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	err error

	state state.State
	ctx   data.MSGenContext

	stateMachine *state.StateMachine[data.MSGenContext]
}

func NewModel(args *data.MSGenArgs) (*Model, error) {
	sm := state.NewStateMachine[data.MSGenContext]()

	if err := registerStates(sm); err != nil {
		return nil, err
	}

	if err := registerTransitions(sm); err != nil {
		return nil, err
	}

	return &Model{
		state:        state.StateNone,
		stateMachine: sm,
		ctx: data.MSGenContext{
			Args: args,
			Data: data.NewMSGen(),
		},
	}, nil
}

func (m *Model) Init() tea.Cmd {
	return m.TransitionTo(boot.State)
}

func (m *Model) Quit(err error) tea.Cmd {
	m.err = err
	return tea.Quit
}

func (m *Model) Error() error {
	return m.err
}

func (m *Model) TransitionTo(next state.State) tea.Cmd {
	if next == state.StateNone {
		return m.Quit(nil)
	}

	if !m.stateMachine.HasState(next) {
		return m.Quit(state.NewStateError(next, "state not found"))
	}

	m.state = next
	return state.SendMsg(state.NewEnterMsg(next))
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, m.Quit(nil)
		}

	case state.StateMsg:
		if msg.State() != m.state {
			str := fmt.Sprintf("state msg mismatch: expected %s, got %s", m.state, msg.State())
			return m, m.Quit(state.NewStateError(m.state, str))
		}

		switch msg := msg.(type) {
		case state.ErrorMsg:
			return m, m.Quit(msg.Error())

		case state.CompletionMsg:
			next := m.stateMachine.Next(m.state)
			return m, m.TransitionTo(next)
		}
	}

	return m, m.stateMachine.Update(m.state, &m.ctx, msg)
}

func (m *Model) View() string {
	if !m.stateMachine.HasState(m.state) {
		// Developer panic: ensure all states are registered in the state machine
		panic("unrenderable state: " + string(m.state))
	}

	output := m.stateMachine.View(m.state, &m.ctx)

	if !strings.HasSuffix(output, "\n") {
		output += "\n"
	}

	return output
}

func registerStates(sm *state.StateMachine[data.MSGenContext]) error {
	if err := sm.AddState(boot.State, boot.Update, boot.View); err != nil {
		return err
	}
	return nil
}

func registerTransitions(sm *state.StateMachine[data.MSGenContext]) error {
	return nil
}
