package v1

import (
	//"github.com/charmbracelet/bubbles/key"

	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// #TODO adding items to list

var (
	editors = []string{"vim", "nvim", "code", "subl"}
	editor  string
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
	// Data
	data Data
)

func removeByIndex[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

type Data struct {
	Item   []Item `json:"item"`
	Editor string `json:"editor"`
}

type Item struct {
	TITLE string `json:"title"`
	DESC  string `json:"desc"`
}

func (i Item) Title() string       { return i.TITLE }
func (i Item) Description() string { return i.DESC }
func (i Item) FilterValue() string { return i.DESC }

// Models

type State int

const (
	ItemList     State = iota
	EditorChoice
	TextInput
)

type model struct {
	ListModel  ListModel
	InputModel InputModel
	Editor     textinput.Model
	state      State
}

type ListModel struct {
	list list.Model
}
type EditorListModel struct {
	list list.Model
}

type InputModel struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode textinput.CursorMode
}

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

func initialModel() model {

	data = data.GetData()
	editor = data.Editor
	var items []list.Item
	for i := 0; i < len(data.Item); i++ {

		items = append(items, data.Item[i])

	}

	delegate := list.NewDefaultDelegate()

	m := model{ListModel{list: list.New(items, delegate, 0, 0)},
		InputModel{inputs: make([]textinput.Model, 2)},
		textinput.New(), State(ItemList)}

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

	m.Editor.Placeholder = "0"
	m.Editor.PromptStyle = focusedStyle
	m.Editor.TextStyle = focusedStyle
	m.Editor.CharLimit = 1

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func ListUpdate(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			choice := m.ListModel.list.Items()[m.ListModel.list.Index()].FilterValue()
			OpenFile(choice, editor)
			return m, tea.Quit
		case "D": //remove item
			removeItem(data.Item[m.ListModel.list.Index()].FilterValue())
			m.ListModel.list.RemoveItem(m.ListModel.list.Index())
			return m, nil
		case "A": // add item
			m.state = TextInput
			return m, nil
		case "E": // change editor
			m.state = EditorChoice
			return m, nil
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.ListModel.list.SetSize(msg.Width-h, msg.Height-v)
	}
	m.ListModel.list, cmd = m.ListModel.list.Update(msg)
	return m, cmd
}

func InputUpdate(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	input := m.InputModel

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
                        m.state = ItemList
			return m, nil

			//change cursor Mode
		case "ctrl+r":
			input.cursorMode++
			if input.cursorMode > textinput.CursorHide {
				input.cursorMode = textinput.CursorBlink
			}
			cmds := make([]tea.Cmd, len(input.inputs))
			for i := range input.inputs {
				cmds[i] = input.inputs[i].SetCursorMode(input.cursorMode)
			}
			m.InputModel = input
			return m, tea.Batch(cmds...)

			//set focus to next input
		case "enter", "up", "down", "shift+tab", "tab":
			s := msg.String()

			if s == "enter" && input.focusIndex == len(input.inputs) {
				if input.inputs[0].Value() != "" && input.inputs[1].Value() != "" {
					title := input.inputs[0].Value()
					path := input.inputs[1].Value()
					addItem(title, path)
					cmd = m.ListModel.list.InsertItem(len(data.Item), Item{TITLE: title, DESC: path})

					m.InputModel = input
					m.state = ItemList

					input.inputs[0].SetValue("")
					input.inputs[1].SetValue("")
					return m, cmd
				}
			}

			// cycle indexes
			if s == "up" || s == "shift+tab" {
				input.focusIndex--
			} else {
				input.focusIndex++
			}

			if input.focusIndex > len(input.inputs) {
				input.focusIndex = 0
			} else if input.focusIndex < 0 {
				input.focusIndex = len(input.inputs)
			}

			cmds := make([]tea.Cmd, len(input.inputs))
			for i := 0; i <= len(input.inputs)-1; i++ {
				if i == input.focusIndex {
					// set focusedStyle
					cmds[i] = input.inputs[i].Focus()
					input.inputs[i].PromptStyle = focusedStyle
					input.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focusedStyle
				input.inputs[i].Blur()
				input.inputs[i].PromptStyle = noStyle
				input.inputs[i].TextStyle = noStyle
			}
			m.InputModel = input
			return m, tea.Batch(cmds...)
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.ListModel.list.SetSize(msg.Width-h, msg.Height-v)
	}

	cmd = m.updateInputs(msg)

	return m, cmd
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.InputModel.inputs))

	for i := range m.InputModel.inputs {
		m.InputModel.inputs[i], cmds[i] = m.InputModel.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func EditorChoiceUpdate(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

        m.Editor.Focus()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			index, _ := strconv.Atoi(m.Editor.Value())
			data.Editor = editors[index]
			saveData(data)
                        data = data.GetData()
			m.state = ItemList
			return m, nil
		case "cntl+l":
			m.state = ItemList
			return m, nil
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.ListModel.list.SetSize(msg.Width-h, msg.Height-v)
	}

	m.Editor, cmd = m.Editor.Update(msg)

	return m, cmd
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

func renderInputs(m model) string {
	var b strings.Builder
	input := m.InputModel

	for i := range input.inputs {
		b.WriteString(input.inputs[i].View())
		if i < len(input.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if input.focusIndex == len(input.inputs) {
		button = &focusedButton
	}

	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(helpStyle.Render("cursor Mode is "))
	b.WriteString(cursorModeHelpStyle.Render(input.cursorMode.String()))
	b.WriteString(helpStyle.Render("(ctrl+r to change style"))

	return b.String()
}

func renderEditor(m model) string {
	var b strings.Builder

	b.WriteString("Choose your editor of choice.\n")
	for i := 0; i < len(editors); i++ {
		fmt.Fprintf(&b, "\n%v.%s\n", i, editors[i])
	}

	fmt.Fprintf(&b, "\n%s\n", m.Editor.View())

	return b.String()
}

func (m model) View() string {
	if m.state == EditorChoice {
		return docStyle.Render(renderEditor(m))
	}
	if m.state == TextInput {
		return docStyle.Render(renderInputs(m))
	}
        ed := helpStyle.Render("\nEditor(" + data.Editor + ")")
        
        return docStyle.Render(m.ListModel.list.View()) + ed
}

func Start() error {

	m := initialModel()

	p := tea.NewProgram(m, tea.WithAltScreen())
	err := p.Start()
	return err
}
