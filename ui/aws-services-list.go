package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type ListItem struct {
	title, desc string
}

func (li ListItem) Title() string       { return li.title }
func (li ListItem) Description() string { return li.desc }
func (li ListItem) FilterValue() string { return li.title }

type AwsDocModel struct {
	servicesListComponent list.Model // bubbletea list component
	services              []string   // available aws services to get help
	cursor                int        // current cursor position
	selected              int        // selected aws service
}

func (aws AwsDocModel) Init() tea.Cmd {
	return nil
}

func (aws AwsDocModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		aws.servicesListComponent.SetSize(msg.Width-h, msg.Height-v)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return aws, tea.Quit
		case "down", "j":
			if aws.cursor < len(aws.services)-1 {
				aws.cursor++
			}
		case "up", "k":
			if aws.cursor > 0 {
				aws.cursor--
			}
		case "enter":
			aws.selected = aws.cursor
		}
	}

	var cmd tea.Cmd
	aws.servicesListComponent, cmd = aws.servicesListComponent.Update(msg)
	return aws, cmd
}

func (aws AwsDocModel) View() string {
	return docStyle.Render(aws.servicesListComponent.View())
}

func ServicesListModel() tea.Model {
	l := []list.Item{
		ListItem{title: "S3", desc: "Documentation resources for aws simple storage service."},
		ListItem{title: "DynamoDB", desc: "Documentation resources for aws dynamo db service internals and api."},
		ListItem{title: "Lambda", desc: "Documentation resources for aws serverles lambda functions."},
		ListItem{title: "EC2", desc: "Documentation resources for aws EC2 api, configuration, and internals"},
	}

	m := AwsDocModel{servicesListComponent: list.New(l, list.NewDefaultDelegate(), 0, 0)}
	m.servicesListComponent.Title = "Available AWS Documentation Resources"

	return m
}
