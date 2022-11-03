package v1

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

func path(dev *bool) string{
    if (*dev) {
        // change the data file for development & don't change the data.json
        // template so the install.sh still works
        return "./dev_data.json"
    }
        user, err := user.Current()
        if err != nil {
                log.Fatal(err)
        }
        path := fmt.Sprintf("/home/%s/openup/data.json", user.Username)
        return path
}

func  (d Data) GetData() Data{
        file, err := ioutil.ReadFile(path(dev))
	if err != nil {
                log.Fatalf("failed to open the file: %e", err)
	}
        
        err = json.Unmarshal([]byte(file), &d)

        if err != nil {
                log.Fatalf("failed parsing json: %e", err)
        }

	return d 
}

// SetDefaultEditor sets the default editor if it exists as an environment 
// variable. If it does not exists nothing happens to the data struct.
func (d *Data) SetDefaultEditor() {
    defaultEditor, success := os.LookupEnv("EDITOR")

    if (success && defaultEditor != "") {
        d.Editor = defaultEditor
    }
}

func saveData(d Data) {
        parsed, err := json.Marshal(d)
        if err != nil {
                log.Fatal(err)
        }

        err = ioutil.WriteFile(path(dev), parsed, fs.FileMode(0))
        if err != nil {
                log.Fatal(err)
        }
}
