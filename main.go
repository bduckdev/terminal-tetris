package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalln("err: ", err)
	}
	defer f.Close()

	p := tea.NewProgram(InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Something went bad", err)
		os.Exit(1)
	}
}
