package v1

import (
	"github.com/charmbracelet/bubbles/key"
)

// all the keybindings
type ListKeyMap struct {
	deleteItem   key.Binding
	addItem      key.Binding
	changeEditor key.Binding
}

// holds the key as a char (string) and the desc as a string
type KeyMap struct {
	Key  string
	Desc string
}

// KeyMaps struct
type KeyMaps struct {
	DeleteItem   KeyMap
	AddItem      KeyMap
	ChangeEditor KeyMap
}

// TODO: deserialisation function for keymaps
func newListKeyMap(keyMaps KeyMaps) *ListKeyMap {
	return &ListKeyMap{
		deleteItem: key.NewBinding(
			key.WithKeys(keyMaps.DeleteItem.Key),
			key.WithHelp(keyMaps.DeleteItem.Key, keyMaps.DeleteItem.Desc),
		),
		addItem: key.NewBinding(
			key.WithKeys(keyMaps.AddItem.Key),
			key.WithHelp(keyMaps.AddItem.Key, keyMaps.AddItem.Desc),
		),
		changeEditor: key.NewBinding(
			key.WithKeys(keyMaps.ChangeEditor.Key),
			key.WithHelp(keyMaps.ChangeEditor.Key, keyMaps.ChangeEditor.Desc),
		),
	}
}
