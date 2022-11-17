package v1

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func OpenFile(path string, editor string, cdToDir bool) {
	cmd := exec.Command(editor, path)
        if cdToDir {
                os.(filepath.Dir(path))
        }
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
