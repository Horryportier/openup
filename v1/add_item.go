package v1

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)


type InputModel struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode textinput.CursorMode
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
					if changeItem {
						changeItem = false
						removeItem(changetmp.DESC)
						m.ListModel.list.RemoveItem(m.ListModel.list.Index())
					}
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
		h, v := dockstyle.GetFrameSize()
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
