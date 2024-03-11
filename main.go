package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("36")
	s.InputField = lipgloss.NewStyle().BorderForeground(s.BorderColor).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(40)
	return s
}

type Main struct {
	index     int
	done      bool
	questions []Question
	width     int
	height    int
	styles    *Styles
}

type Question struct {
	question string
	answer   string
	input    Input
}

func newQuestion(question string) Question {
	return Question{question: question}
}

func newShortQuestion(q string) Question {
	question := newQuestion(q)
	model := NewShortAnswerField()
	question.input = model
	return question
}

func newLongQuestion(q string) Question {
	question := newQuestion(q)
	model := NewLongAnswerField()
	question.input = model
	return question
}

func New(questions []Question) *Main {
	styles := DefaultStyles()
	return &Main{
		styles:    styles,
		questions: questions,
	}
}

func (m Main) Init() tea.Cmd {
	return nil
}

func (m Main) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	current := &m.questions[m.index]

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if m.index == len(m.questions)-1 {
				m.done = true
			} else if m.index == len(m.questions) {
				return m, tea.Quit
			}

			current.answer = current.input.Value()
			m.Next()
			return m, current.input.Blur
		}
	}
	current.input, cmd = current.input.Update(msg)
	return m, cmd
}

func (m Main) View() string {
	current := m.questions[m.index]
	if m.done {
		var output string
		for _, q := range m.questions {
			output += q.question + " " + q.answer + "\n"
		}
		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			output,
		)
	}
	if m.width == 0 || m.height == 0 {
		return "Initializing..."
	}
	return lipgloss.Place(
		m.width,
		m.height,

		lipgloss.Center,
		lipgloss.Center,

		lipgloss.JoinVertical(
			lipgloss.Center,
			m.questions[m.index].question,
			m.styles.InputField.Render(current.input.View()),
		),
	)
}

func (m *Main) Next() {
	if m.index < len(m.questions)-1 {
		m.index++
	} else {
		m.index = 0
	}
}

func main() {
	questions := []Question{
		newShortQuestion("What is your name?"),
		newLongQuestion("What is your quest?"),
		newShortQuestion("What is your favorite color?"),
	}

	m := New(questions)

	f, err := tea.LogToFile("log.txt", "debug")
	if err != nil {
		log.Fatal("Err %w", err)
	}
	defer f.Close()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal("Err %w", err)
	}
}
