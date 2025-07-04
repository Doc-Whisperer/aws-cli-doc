package main

import (
	"fmt"
	"os"

	"doc-whisperer/aws-doc-agent-cli/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(ui.InitialAwsDocModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("An error occurred: %f", err)
		os.Exit(1)
	}
}
