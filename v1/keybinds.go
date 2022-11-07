package v1

import (
    "github.com/charmbracelet/bubbles/key"
)

// all the keybindings
type ListKeyMap struct {
    deleteItem key.Binding
    addItem key.Binding
    changeEditor key.Binding
}

// all the key functions
func newListKeyMap() *ListKeyMap {
	return &ListKeyMap{
        deleteItem: key.NewBinding(
        key.WithKeys("D"),
        key.WithHelp("D", "remove an item"),
            ),
        addItem: key.NewBinding(
        key.WithKeys("A"),
        key.WithHelp("A", "add an item"),
            ),
        changeEditor: key.NewBinding(
        key.WithKeys("E"),
        key.WithHelp("E", "change the editor"),
            ),
    }
}
