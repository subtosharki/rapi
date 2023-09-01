package lib

import "os"

func Info(message string) {
	println("RAPI INFO: " + message)
}

func Error(error string) {
	println("RAPI ERROR: " + error)
	os.Exit(1)
}
