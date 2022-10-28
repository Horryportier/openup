package v1

import (
	"log"
	"os"
	"os/exec"
)


func OpenFile(path string, editor string)  {
        cmd := exec.Command(editor, path)
        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        err := cmd.Run()
        if err != nil {
                log.Fatal(err)
        }
}
