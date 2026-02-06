package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gnujesus/gnu-finance/internal/data"
)

// In order for the model to work, it requires 3 methods:
// Init, Update and View

type Model struct {
	Choices []string // the options list
	Cursor  int      // which item am i currently pointing at
	Company data.CompanyInfo
}

func InitialModel(c data.CompanyInfo) Model {
	return Model{
		Choices: []string{"Simple View", "Detailed View", "Graphs", "Exit"},
		Company: c,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}

		case "down", "j":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}

		case "enter", " ":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	// header
	s := "Select an option\n\n"

	for i, choice := range m.Choices {
		Cursor := " "

		if m.Cursor == i {
			Cursor = ">"
		}

		s += fmt.Sprintf("%s [%s]\n", Cursor, choice)
	}

	s += "\nPress q to quit. \n"

	return s
}
