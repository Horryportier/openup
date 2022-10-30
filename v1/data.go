package v1

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
)


func  (d Data) GetData() Data{
	file, err := ioutil.ReadFile("v1/file.json")
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

        err = ioutil.WriteFile("v1/file.json", parsed, fs.FileMode(0))
        if err != nil {
                log.Fatal(err)
        }
}
