package main

import (
	"fmt"
	"github.com/subtosharki/rapi/src/cmd"
)

func main() {
	err := cmd.Root.Execute()
	if err != nil {
		fmt.Printf("RAPI Command Error: %s", err)
	}
}
