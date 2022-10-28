package v1

import (
	//"github.com/charmbracelet/bubbles/key"

	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// #TODO add text input.
// KEYBINDS:
// -standard list bindings
// -change existing item {c}
// -delete item {D}
// -add item {a}
// -change editor {cntl+e}

var (
	editor string
	// list style
	docStyle = lipgloss.NewStyle().Margin(1, 2)

	// text imput style
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("2O5"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle.Copy()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Copy().Render("[Add item]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Add item"))
)

type Data struct {
	Item   []Item
	Editor string
}

type Item struct {
	TITLE string
	DESC  string
}

func (i Item) Title() string       { return i.TITLE }
func (i Item) Description() string { return i.DESC }
func (i Item) FilterValue() string { return i.DESC }

// Models

type model struct {
	ListModel  ListModel
	InputModel InputModel
}

type ListModel struct {
	list list.Model
}

type InputModel struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode textinput.CursorMode
}

func initialModel() model {

	var data Data
	data = data.GetData()
	editor = data.Editor

	var items []list.Item
	for i := 0; i < len(data.Item); i++ {

		items = append(items, data.Item[i])

	}

	m := model{ListModel{list: list.New(items, list.NewDefaultDelegate(), 0, 0)},
		InputModel{inputs: make([]textinput.Model, 2)}}
	m.ListModel.list.Title = "FBI OPEN UP"

        var t textinput.Model
        for i := range m.InputModel.inputs {
                t = textinput.New()
                t.CursorStyle = cursorStyle
                t.CharLimit = 128

                switch i {
                case 0:
                        t.Placeholder = "Title"
                        t.Focus()
                        t.PromptStyle = focusedStyle
                        t.TextStyle = focusedStyle
                case 1:
                        t.Placeholder = "Path"
                        t.CharLimit = 256
                }
                m.InputModel.inputs[i] = t
        }

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

// TODO:add text input
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			choice := m.ListModel.list.Items()[m.ListModel.list.Index()].FilterValue()
			OpenFile(choice, editor)
			return m, tea.Quit
		}
		if msg.String() == "D" {
			m.ListModel.list.RemoveItem(m.ListModel.list.Index())
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.ListModel.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.ListModel.list, cmd = m.ListModel.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.ListModel.list.View())
}

func Start() error {

	m := initialModel()

	p := tea.NewProgram(m, tea.WithAltScreen())
	err := p.Start()
	return err
}
