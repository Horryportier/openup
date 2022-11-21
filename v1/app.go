package v1

import (
	"flag"
	// "log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// #TODO adding items to list

var (
	editors = []string{"vim", "nvim", "code", "subl"}
	editor  string

	// Data
	data Data
)

type Data struct {
	Item    []Item  `json:"item"`
	Editor  string  `json:"editor"`
	KeyMaps KeyMaps `json:"keymaps"`
}

type State int

const (
	ItemList State = iota
	EditorChoice
	TextInput
)

type model struct {
	ListModel  ListModel
	InputModel InputModel
	Editor     textinput.Model
	keys       *ListKeyMap
	state      State
}

func initialModel() model {

	flag.Parse()
	data = data.GetData()
	keymaps := GetConfig()
	if !*noDefaultEditor {
		data.SetDefaultEditor()
	}
	editor = data.Editor
	var (
		items    []list.Item
		listKeys = newListKeyMap(keymaps)
	)
	for i := 0; i < len(data.Item); i++ {

		items = append(items, data.Item[i])

	}

	delegate := list.NewDefaultDelegate()

	m := model{ListModel{list: list.New(items, delegate, 0, 0)},
		InputModel{inputs: make([]textinput.Model, 2)},
		textinput.New(), listKeys, State(ItemList)}

	m.ListModel.list.Title = "FBI OPEN UP"

	m.ListModel.list.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.addItem,
			listKeys.changeEditor,
			listKeys.deleteItem,
		}
	}

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

	m.Editor.Placeholder = "0"
	m.Editor.PromptStyle = focusedStyle
	m.Editor.TextStyle = focusedStyle
	m.Editor.CharLimit = 1

        m = SetStyle(m)

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if m.state == ItemList {
		return ListUpdate(m, msg)
	}
	if m.state == TextInput {
		return InputUpdate(m, msg)
	}
	if m.state == EditorChoice {
		return EditorChoiceUpdate(m, msg)
	}

	return m, nil
}

func (m model) View() string {
	if m.state == EditorChoice {
		return dockstyle.Render(renderEditor(m))
	}
	if m.state == TextInput {
		return dockstyle.Render(renderInputs(m))
	}
	ed := helpStyle.Render("\nEditor(" + accentColor1Text.Render(data.Editor) + ")")

	return dockstyle.Render(m.ListModel.list.View() + ed)
}

func Start() error {

	m := initialModel()

	p := tea.NewProgram(m, tea.WithAltScreen())
	err := p.Start()
	return err
}
