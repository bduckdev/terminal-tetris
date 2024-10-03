package main

import "github.com/charmbracelet/lipgloss"

type styles struct {
	Program lipgloss.Style
}

func DefaultStyles() styles {
	return styles{
		Program: lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0),
	}
}
