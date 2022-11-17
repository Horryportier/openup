package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"

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

// get the path to config file
func configPath(dev *bool) string {
	if *dev {
		// change the data file for development & don't change the data.json
		// template so the install.sh still works
		return "./dev_config.json"
	}
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	path := fmt.Sprintf("/home/%s/.config/openup/config.json", user.Username)
	return path
}

// Unmarshal Keymaps in case of more configurations one could extract the
// configPath and the GetConfig
func GetConfig() KeyMaps {
	var keymaps KeyMaps
	file, err := ioutil.ReadFile(configPath(dev))
	if err != nil {
        return keymaps
	}

	err = json.Unmarshal([]byte(file), &keymaps)

	if err != nil {
		log.Fatalf("failed parsing keymaps json: %e", err)
	}
	return keymaps
}

// initialize keymaps from data.json file
// if data.json file has no keymaps the default keymaps get set
func newListKeyMap(keyMaps KeyMaps) *ListKeyMap {
	if keyMaps.DeleteItem.Key == "" {
		keyMaps.DeleteItem.Key = "D"
		keyMaps.DeleteItem.Desc = "delete an item"
	}
	if keyMaps.ChangeEditor.Key == "" {
		keyMaps.ChangeEditor.Key = "E"
		keyMaps.ChangeEditor.Desc = "change the editor"
	}
	if keyMaps.AddItem.Key == "" {
		keyMaps.AddItem.Key = "A"
		keyMaps.AddItem.Desc = "add an item"
	}
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
