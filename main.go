package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type AwsDocModel struct {
	services map[int]string // available aws services to get help
	service  int            // selected aws service
	prompt   string         // prompt for aws doc
}

func InitialAwsDocModel() AwsDocModel {
	return AwsDocModel{
		services: map[int]string{
			1: "s3",
			2: "lambda",
			3: "dynamoDb",
			4: "ec2",
		},
		service: 1,
		prompt:  "",
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
			if adm.service < len(adm.services)-1 {
				adm.service++
			}
		case "up", "k":
			if adm.service > 1 {
				adm.service--
			}
		case "enter":
			adm.prompt = adm.services[adm.service]
		}
	}

	return adm, nil
}

func (adm AwsDocModel) View() string {
	s := "Which AWS doc do you need help with?\n\n"

	for i, service := range adm.services {
		cursor := " "
		if adm.service == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s [%s]\n", cursor, service)
	}

	s += "\npress q to quit.\n"

	return s
}

func main() {
	p := tea.NewProgram(InitialAwsDocModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("An error occurred: %f", err)
		os.Exit(1)
	}
}
