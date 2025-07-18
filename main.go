package main

import (
	"fmt"
	"os"

	"doc-whisperer/aws-doc-agent-cli/ui/list"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(list.ServicesListModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("An error occurred: %f", err)
		os.Exit(1)
	}
}
