package v1

import (
	"encoding/json"
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
