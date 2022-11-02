package v1

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os/user"
)

func path() string{
        user, err := user.Current()
        if err != nil {
                log.Fatal(err)
        }
        path := fmt.Sprintf("/home/%s/openup/data.json", user.Username)
        return path
}

func  (d Data) GetData() Data{
        file, err := ioutil.ReadFile(path())
	if err != nil {
                log.Fatalf("failed to open the file: %e", err)
	}
        
        err = json.Unmarshal([]byte(file), &d)

        if err != nil {
                log.Fatalf("failed parsing json: %e", err)
        }

	return d 
}

func saveData(d Data) {
        parsed, err := json.Marshal(d)
        if err != nil {
                log.Fatal(err)
        }

        err = ioutil.WriteFile(path(), parsed, fs.FileMode(0))
        if err != nil {
                log.Fatal(err)
        }
}
