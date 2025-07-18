package prompt

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type PromptModel struct {
	value   string
	service string
}

func RenderModel(service string) tea.Model {
	return PromptModel{
		value:   "Rendering Prompt Model",
		service: service,
	}
}

func (p PromptModel) Init() tea.Cmd {
	return nil
}

func (p PromptModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}

	return p, nil
}

func (p PromptModel) View() string {
	s := "\nInside the prompt TUI"
	s += fmt.Sprintf("\n\nSelected service: %s", p.service)
	return s
}
