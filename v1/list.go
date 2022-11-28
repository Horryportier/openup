package v1

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ListModel struct {
	list list.Model
}

type Item struct {
	TITLE string `json:"title"`
	DESC  string `json:"desc"`
}

func (i Item) Title() string       { return i.TITLE }
func (i Item) Description() string { return i.DESC }
func (i Item) FilterValue() string { return i.DESC }
func addItem(title string, path string) {
	data.Item = append(data.Item, Item{TITLE: title, DESC: path})
	saveData(data)
}

func removeItem(desc string) {
	for i := 0; i < len(data.Item); i++ {
		if data.Item[i].DESC == desc {
			data.Item = removeByIndex(data.Item, i)
		}
	}
	saveData(data)
}

func ListUpdate(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			if len(m.ListModel.list.Items()) > 0 {
				choice := m.ListModel.list.Items()[m.ListModel.list.Index()].FilterValue()
				OpenFile(choice, editor)
				return m, tea.Quit
			}
		}
		switch {
		case key.Matches(msg, m.keys.deleteItem): //remove item
			if len(m.ListModel.list.Items()) > 0 {
				removeItem(data.Item[m.ListModel.list.Index()].FilterValue())
				m.ListModel.list.RemoveItem(m.ListModel.list.Index())
			}
			return m, nil
		case key.Matches(msg, m.keys.addItem): // add item
			m.state = TextInput
			return m, nil
		case key.Matches(msg, m.keys.changeEditor): // change editor
			m.state = EditorChoice
			return m, nil
		case key.Matches(msg, m.keys.changeItem):
			i := m.ListModel.list.Index()
			m.InputModel.inputs[0].SetValue(data.Item[i].TITLE)
			m.InputModel.inputs[1].SetValue(data.Item[i].DESC)
                        changeItem = true
			m.state = TextInput
			return m, nil
		}

	case tea.WindowSizeMsg:
		h, v := dockstyle.GetFrameSize()
		m.ListModel.list.SetSize(msg.Width-h, msg.Height-v)
		dockstyle.Width(msg.Width)
		dockstyle.Height(msg.Height)
	}

	m.ListModel.list, cmd = m.ListModel.list.Update(msg)
	return m, cmd
}
