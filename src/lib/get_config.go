package lib

import "github.com/spf13/viper"

type Config struct {
	Framework       string `json:"framework"`
	RoutesPath      string `json:"routespath"`
	MiddlewaresPath string `json:"middlewarespath"`
	MainFilePath    string `json:"mainfile"`
	ProjectName     string `json:"projectname"`
}

func GetConfig() Config {
	LoadConfig()

	framework := viper.GetString("framework")
	if framework == "" {
		Error("framework not found in rapi.json")
		ExitBad()
	}
	routesPath := viper.GetString("routespath")
	if routesPath == "" {
		Error("routespath not found in rapi.json")
		ExitBad()
	}
	middlewaresPath := viper.GetString("middlewarespath")
	if middlewaresPath == "" {
		Error("middlewarespath not found in rapi.json")
		ExitBad()
	}
	mainFile := viper.GetString("mainfilepath")
	if mainFile == "" {
		Error("mainfilepath not found in rapi.json")
		ExitBad()
	}
	projectName := viper.GetString("projectname")
	if projectName == "" {
		Error("projectname not found in rapi.json")
		ExitBad()
	}
	config := Config{
		Framework:       framework,
		RoutesPath:      routesPath,
		MiddlewaresPath: middlewaresPath,
		MainFilePath:    mainFile,
		ProjectName:     projectName,
	}
	VerifyConfig(config)
	return config
}
