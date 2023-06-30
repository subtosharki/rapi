package lib

import "github.com/spf13/viper"

func LoadConfig() {
	viper.SetConfigName("rapi")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		Error("Error reading config file")
		ExitBad()
	}
}
