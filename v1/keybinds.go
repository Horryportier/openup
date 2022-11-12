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

// holds the key as a char (string) and the desc as a string
type KeyMap struct {
    key string
    desc string
}

// KeyMaps struct
type KeyMaps struct {
    deleteItem KeyMap
    addItem KeyMap
    changeEditor KeyMap
}

// TODO: deserialisation function for keymaps
func newListKeyMap(keyMaps KeyMaps) *ListKeyMap {
	return &ListKeyMap{
        deleteItem: key.NewBinding(
        key.WithKeys(keyMaps.deleteItem.key),
        key.WithHelp(keyMaps.deleteItem.key, keyMaps.deleteItem.desc),
            ),
        addItem: key.NewBinding(
        key.WithKeys(keyMaps.addItem.key),
        key.WithHelp(keyMaps.addItem.key, keyMaps.addItem.desc),
            ),
        changeEditor: key.NewBinding(
        key.WithKeys(keyMaps.changeEditor.key),
        key.WithHelp(keyMaps.changeEditor.key, keyMaps.changeEditor.desc),
            ),
    }
}

