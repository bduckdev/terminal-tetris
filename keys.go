package main

import "github.com/charmbracelet/bubbles/key"

type keymap struct {
	quit key.Binding
}

func DefaultMap() keymap {
	return keymap{
		quit: key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
	}
}
