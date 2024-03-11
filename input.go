package main

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Input interface {
	Value() string
	Blur() tea.Msg
	View() string
	Update(tea.Msg) (Input, tea.Cmd)
}

/* SHORT ANSWER FIELD */

type ShortAnswerField struct {
	textInput textinput.Model
}

func (sa *ShortAnswerField) Value() string {
	return sa.textInput.Value()
}

func (sa *ShortAnswerField) Blur() tea.Msg {
	return sa.textInput.Blur
}

func (sa *ShortAnswerField) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	sa.textInput, cmd = sa.textInput.Update(msg)
	return sa, cmd
}

func (sa *ShortAnswerField) View() string {
	return sa.textInput.View()
}

/* LONG ANSWER FIELD */

type LongAnswerField struct {
	textArea textarea.Model
}

// textInput
func NewShortAnswerField() *ShortAnswerField {
	ti := textinput.New()
	ti.Focus()
	ti.Placeholder = "Type your answer here"
	return &ShortAnswerField{ti}
}

func (la *LongAnswerField) Value() string {
	return la.textArea.Value()
}

func (la *LongAnswerField) Blur() tea.Msg {
	return la.textArea.Blur
}

func (la *LongAnswerField) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	la.textArea, cmd = la.textArea.Update(msg)
	return la, cmd
}

// textArea
func NewLongAnswerField() *LongAnswerField {
	ta := textarea.New()
	ta.Focus()
	ta.Placeholder = "Type your long answer here"
	return &LongAnswerField{ta}
}

func (la *LongAnswerField) View() string {
	return la.textArea.View()
}
