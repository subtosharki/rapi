package lib

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Framework       string `json:"framework"`
	RoutesPath      string `json:"routespath"`
	MiddlewaresPath string `json:"middlewarespath"`
	MainFilePath    string `json:"mainfile"`
	ProjectName     string `json:"projectname"`
}

func (c *Config) Get() {
	_, err := os.Stat("rapi.json")
	if err != nil {
		Error("rapi.json not found, please run rapi init")
	}
	viper.SetConfigName("rapi")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		Error("Error reading config file")
	}

	framework := viper.GetString("framework")
	if framework == "" {
		Error("framework not found in rapi.json")
	}
	routesPath := viper.GetString("routespath")
	if routesPath == "" {
		Error("routespath not found in rapi.json")
	}
	middlewaresPath := viper.GetString("middlewarespath")
	if middlewaresPath == "" {
		Error("middlewarespath not found in rapi.json")
	}
	mainFile := viper.GetString("mainfilepath")
	if mainFile == "" {
		Error("mainfilepath not found in rapi.json")
	}
	projectName := viper.GetString("projectname")
	if projectName == "" {
		Error("projectname not found in rapi.json")
	}
	config := Config{
		Framework:       framework,
		RoutesPath:      routesPath,
		MiddlewaresPath: middlewaresPath,
		MainFilePath:    mainFile,
		ProjectName:     projectName,
	}
	_, err = os.Stat(config.RoutesPath)
	if err != nil {
		Error("Invalid path to routes")
	}
	_, err = os.Stat(config.MiddlewaresPath)
	if err != nil {
		Error("Invalid path to middlewares")
	}
	_, err = os.Stat(config.MainFilePath)
	if err != nil {
		Error("Invalid path to main.go")
	}
	switch config.Framework {
	case "gin":
	case "echo":
	case "fiber":
	case "chi":
	default:
		Error("Invalid framework")
	}
	*c = config
}

func (_ *Config) Setup(config Config) {
	viper.AddConfigPath(".")
	viper.SetConfigName("rapi")
	viper.SetConfigType("json")
	viper.Set("projectname", config.ProjectName)
	viper.Set("framework", config.Framework)
	viper.Set("routespath", config.RoutesPath)
	viper.Set("middlewarespath", config.MiddlewaresPath)
	viper.Set("mainfilepath", config.MainFilePath)
	err := viper.SafeWriteConfig()
	ErrorCheck(err)
}
