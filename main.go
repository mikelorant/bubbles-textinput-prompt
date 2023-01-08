package main

import (
	"log"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type model struct {
	textInputA textinput.Model
	textInputB textinput.Model
}

func initialModel() model {
	tiA := makeTextInput("1234")
	tiB := makeTextInput("1234567890")

	return model{
		textInputA: tiA,
		textInputB: tiB,
	}
}

func makeTextInput(str string) textinput.Model {
	ti := textinput.New()
	ti.Prompt = str
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return ti
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}

	m.textInputA, cmd = m.textInputA.Update(msg)
	cmds = append(cmds, cmd)
	m.textInputB, cmd = m.textInputB.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	styleA := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Render(m.textInputA.View())

	styleB := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Render(m.textInputB.View())

	both := lipgloss.JoinVertical(lipgloss.Top, styleA, styleB)

	return both
}
