package lib

import "os"

func VerifyConfig(config Config) {
	_, err := os.Stat(config.RoutesPath)
	if err != nil {
		Error("Invalid path to routes")
		ExitBad()
	}
	_, err = os.Stat(config.MiddlewaresPath)
	if err != nil {
		Error("Invalid path to middlewares")
		ExitBad()
	}
	_, err = os.Stat(config.MainFilePath)
	if err != nil {
		Error("Invalid path to main.go")
		ExitBad()
	}
	switch config.Framework {
	case "gin":
	case "echo":
	case "fiber":
	case "chi":
	default:
		Error("Invalid framework")
		ExitBad()
	}
	if config.ProjectName == "" {
		Error("Invalid project name")
		ExitBad()
	}
}
