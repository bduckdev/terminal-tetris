package main

import tea "github.com/charmbracelet/bubbletea"

type model struct {
	playfield [10][20]byte
}

fun InitialModel() model {
	return model{
		playfield: [10][20]byte
	}
}

fun (m Model) Init() tea.Cmd {
	return nil
}

fun (m Model) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return model, nil
}

func (model Model) View() string {
	return ""
}