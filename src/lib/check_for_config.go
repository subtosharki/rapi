package lib

import "os"

func CheckForConfig() {
	_, err := os.Stat("rapi.json")
	if err != nil {
		Error("rapi.json not found, please run rapi init")
		ExitBad()
	}
}
