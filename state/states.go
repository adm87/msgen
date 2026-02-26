package state

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// --------------------------------------------------------------
// State Messages
// --------------------------------------------------------------

type StateMsg interface {
	State() State
}

func SendMsg(msg StateMsg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}

// --------------------------------------------------------------
// State Enter Message
// --------------------------------------------------------------

type EnterMsg struct {
	state State
}

func NewEnterMsg(s State) EnterMsg {
	return EnterMsg{state: s}
}

func (m EnterMsg) State() State {
	return m.state
}

// --------------------------------------------------------------
// State Error Message
// --------------------------------------------------------------

type ErrorMsg struct {
	state State
	err   error
}

func NewErrorMsg(s State, err error) ErrorMsg {
	return ErrorMsg{state: s, err: err}
}

func (m ErrorMsg) State() State {
	return m.state
}

func (m ErrorMsg) Error() error {
	return m.err
}

// --------------------------------------------------------------
// State Completion Message
// --------------------------------------------------------------

type CompletionMsg struct {
	state State
}

func NewCompletionMsg(s State) CompletionMsg {
	return CompletionMsg{state: s}
}

func (m CompletionMsg) State() State {
	return m.state
}

// --------------------------------------------------------------
// State Update Message
// --------------------------------------------------------------

type UpdateMsg struct {
	state State
}

func NewUpdateMsg(s State) UpdateMsg {
	return UpdateMsg{state: s}
}

func (m UpdateMsg) State() State {
	return m.state
}

// --------------------------------------------------------------
// State Errors
// --------------------------------------------------------------

// StateError is an error that occurs when a state encounters an unexpected condition.
type StateError struct {
	State   State  // The state that caused the error
	Message string // The error message
}

// NewStateError creates a new StateError with the given state and message.
func NewStateError(s State, msg string) *StateError {
	return &StateError{State: s, Message: msg}
}

// Error returns the error message for the StateError.
func (e *StateError) Error() string {
	return fmt.Sprintf("state %s: %v", e.State, e.Message)
}

// --------------------------------------------------------------
// State Management
// --------------------------------------------------------------

type State string

const StateNone State = "none"

type StateUpdateFunc[TCtx any] func(ctx *TCtx, msg tea.Msg) tea.Cmd
type StateViewFunc[TCtx any] func(ctx *TCtx) string

// Model represents the current state of the application.
type Model[TCtx any] struct {
	Update StateUpdateFunc[TCtx]
	View   StateViewFunc[TCtx]
}

// StateMachine manages the states and transitions of the application.
type StateMachine[TCtx any] struct {
	states      map[State]Model[TCtx]
	transitions map[State]State
}

func NewStateMachine[TCtx any]() *StateMachine[TCtx] {
	return &StateMachine[TCtx]{
		states:      make(map[State]Model[TCtx]),
		transitions: make(map[State]State),
	}
}

func (sm *StateMachine[TCtx]) AddState(s State, update StateUpdateFunc[TCtx], view StateViewFunc[TCtx]) error {
	if _, exists := sm.states[s]; exists {
		return fmt.Errorf("state %s already exists", s)
	}
	sm.states[s] = Model[TCtx]{
		Update: update,
		View:   view,
	}
	return nil
}

func (sm *StateMachine[TCtx]) AddTransition(from State, to State) error {
	if _, exists := sm.transitions[from]; exists {
		return fmt.Errorf("transition from state %s already exists", from)
	}
	sm.transitions[from] = to
	return nil
}

func (sm *StateMachine[TCtx]) HasState(s State) bool {
	_, exists := sm.states[s]
	return exists
}

func (sm *StateMachine[TCtx]) Next(s State) State {
	if next, ok := sm.transitions[s]; ok {
		return next
	}
	return StateNone
}

func (sm *StateMachine[TCtx]) Update(s State, ctx *TCtx, msg tea.Msg) tea.Cmd {
	if m, ok := sm.states[s]; ok {
		return m.Update(ctx, msg)
	}
	return nil
}

func (sm *StateMachine[TCtx]) View(s State, ctx *TCtx) string {
	if m, ok := sm.states[s]; ok {
		return m.View(ctx)
	}
	return ""
}
