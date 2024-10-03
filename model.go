package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	playfield [10][20]byte
	styles    styles
	keys      keymap
}

func InitialModel() model {
	return model{
		playfield: [10][20]byte{},
		styles:    DefaultStyles(),
		keys:      DefaultMap(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	defaultMap := DefaultMap()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, defaultMap.quit):
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	var output string

	for row := 0; row < 10; row++ {
		for col := range m.playfield[row] {
			switch m.playfield[row][col] {
			case 0:
				output += " "
			case 'I':
			case 'O':
			case 'T':
			case 'S':
			case 'Z':
			case 'J':
			case 'L':
			default:
				output += "?"
			}
		}
		if row < 9 {
			output += "\n"
		}
	}
	return m.styles.Program.Render(output)
}
