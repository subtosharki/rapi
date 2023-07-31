package lib

import (
	"github.com/spf13/viper"
)

func SetupConfig(config Config) {
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
