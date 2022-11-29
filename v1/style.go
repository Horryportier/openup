package v1

import (
	"github.com/charmbracelet/lipgloss"
        "fmt"
)

var (
	primaryColor    = lipgloss.Color("#00FFD2")
	secondaryColor  = lipgloss.Color("#FF4499")
	backgroundColor = lipgloss.Color("#333333")
	accentColor1    = lipgloss.Color("#0a0047")
	accentColor2    = lipgloss.Color("#004687")

	accentColor1Text = lipgloss.NewStyle().Foreground(accentColor1)
	accentColor2Text = lipgloss.NewStyle().Foreground(accentColor2)

	border    = lipgloss.NormalBorder()
	dockstyle = lipgloss.NewStyle().Padding(4).Border(border).Align(lipgloss.Center)

	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("2O5"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle.Copy()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Copy().Render("[Add item]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Add item"))
)

func SetStyle(m model) model {
        //list
	m.ListModel.list.Styles.Title.Foreground(accentColor2)
	m.ListModel.list.Styles.Title.Background(primaryColor)
        m.ListModel.list.Styles.NoItems.Foreground(primaryColor)
        m.ListModel.list.Styles.FilterPrompt.Foreground(primaryColor)
        m.ListModel.list.Styles.HelpStyle.Foreground(primaryColor)
        m.ListModel.list.Styles.FilterCursor.Foreground(secondaryColor)

        for i := range m.InputModel.inputs {
                m.InputModel.inputs[i].TextStyle.Foreground(primaryColor)
                m.InputModel.inputs[i].CursorStyle.Foreground(secondaryColor)
        }



	return m
}
