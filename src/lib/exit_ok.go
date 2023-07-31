package lib

import (
	"os"
	"os/exec"
)

func ExitOk() {
	cmd := exec.Command("go", "fmt", "./...")
	err := cmd.Run()
	ErrorCheck(err)
	os.Exit(0)
}
