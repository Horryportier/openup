package v1

import (
	//"github.com/charmbracelet/bubbles/key"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)
	editor   string
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

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			choice := m.list.Items()[m.list.Index()].FilterValue()
			OpenFile(choice, editor)
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func Start() error {
	var data Data

	data = data.GetData()

	editor = data.Editor

	var items []list.Item
	for i := 0; i < len(data.Item); i++ {

		items = append(items, data.Item[i])

	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "FBI OPEN UP"

	p := tea.NewProgram(m, tea.WithAltScreen())
	err := p.Start()
	return err
}
