package list

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

// Colors
const (
	white       = "#DBD3D3"
	black       = "#000000"
	darkOrange  = "#FF6500"
	lightOrange = "#EC8305"
	darkBlue    = "#0910FF"
	lightBlue   = "#0A4DFF"
	teal        = "#78B9B5"
)

// Constructed Colors
var (
	mainTitleColor = lipgloss.Color(darkOrange)
)

// List Styling
var (
	mainTitleStyling = lipgloss.NewStyle().
				Foreground(mainTitleColor).
				Bold(true)

	selectedTitleColor = lipgloss.
				AdaptiveColor{Light: darkOrange, Dark: lightOrange}
)

var d = list.NewDefaultDelegate()

// Default list delegate with custom styles
func StyledListDelegate() list.DefaultDelegate {
	d.Styles.SelectedTitle = d.Styles.SelectedTitle.Foreground(selectedTitleColor)

	return d
}
