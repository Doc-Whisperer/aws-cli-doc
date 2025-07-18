package list

import (
	"doc-whisperer/aws-doc-agent-cli/ui/prompt"

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
}

func (aws *AwsDocModel) Init() tea.Cmd {
	return nil
}

func (aws *AwsDocModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		aws.servicesListComponent.SetSize(msg.Width-h, msg.Height-v)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return aws, tea.Quit
		case "enter":
			service := aws.servicesListComponent.SelectedItem().FilterValue()
			promptModel := prompt.RenderModel(service)
			return promptModel, nil
		}
	}

	var cmd tea.Cmd
	aws.servicesListComponent, cmd = aws.servicesListComponent.Update(msg)
	return aws, cmd
}

func (aws *AwsDocModel) View() string {
	return docStyle.Render(aws.servicesListComponent.View())
}

func ServicesListModel() tea.Model {
	l := []list.Item{
		ListItem{title: "S3", desc: "Documentation resources for aws simple storage service."},
		ListItem{title: "DynamoDB", desc: "Documentation resources for aws dynamo db service internals and api."},
		ListItem{title: "Lambda", desc: "Documentation resources for aws serverles lambda functions."},
		ListItem{title: "EC2", desc: "Documentation resources for aws EC2 api, configuration, and internals"},
	}

	delegate := StyledListDelegate()

	m := &AwsDocModel{servicesListComponent: list.New(l, delegate, 0, 0)}
	m.servicesListComponent.Title = mainTitleStyling.Render("AWS Services Documentation")

	return m
}
