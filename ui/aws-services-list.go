package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type AwsDocModel struct {
	services []string // available aws services to get help
	cursor   int      // current cursor position
	selected int      // selected aws service
}

func InitialAwsDocModel() AwsDocModel {
	return AwsDocModel{
		services: []string{"DynamoDB", "Lambda", "S3", "EC2"},
	}
}

func (adm AwsDocModel) Init() tea.Cmd {
	return nil
}

func (adm AwsDocModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return adm, tea.Quit
		case "down", "j":
			if adm.cursor < len(adm.services)-1 {
				adm.cursor++
			}
		case "up", "k":
			if adm.cursor > 0 {
				adm.cursor--
			}
		case "enter":
			adm.selected = adm.cursor
		}
	}

	return adm, nil
}

func (adm AwsDocModel) View() string {
	s := "Which AWS doc do you need help with?\n\n"

	for i, service := range adm.services {
		cursor := " "
		if adm.cursor == i {
			cursor = ">"
		}

		checked := " "
		if i == adm.selected {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, service)
	}

	s += "\npress q to quit.\n"

	return s
}
