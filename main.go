package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Initialize the model
	m := NewModel()

	// Create a new Bubble Tea program
	p := tea.NewProgram(m)

	// Start the program
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
