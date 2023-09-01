package lib

import (
	"os"
	"strings"
)

func LoadGoModuleFile() []byte {
	file, err := os.ReadFile("go.mod")
	if err != nil {
		Error("go.mod not found, is this an existing project?")
	}
	if strings.Split(string(file), " ")[0] != "module" {
		Error("Invalid go.mod file")
	}
	return file
}
