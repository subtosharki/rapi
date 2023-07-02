package lib

import (
	"github.com/spf13/viper"
)

func SetupConfig(projectName string, foundFramework string, routesPath string, middlewaresPath string, mainFilePath string) {
	viper.AddConfigPath(".")
	viper.SetConfigName("rapi")
	viper.SetConfigType("json")
	viper.Set("projectname", projectName)
	viper.Set("framework", foundFramework)
	viper.Set("routespath", routesPath)
	viper.Set("middlewarespath", middlewaresPath)
	viper.Set("mainfile", mainFilePath)
	err := viper.SafeWriteConfig()
	ErrorCheck(err)
}
