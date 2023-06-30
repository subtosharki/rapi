package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/subtosharki/rapi/src/lib"
	"os"
	"strings"
)

func init() {
	Root.AddCommand(initCmd)
}

var supportedFrameworks = []string{"fiber", "gin"}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize RAPI in a current project",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat("rapi.json")
		if err == nil {
			lib.RapiError("rapi.json found, already initialized")
			lib.RapiExitBad()
		}
		_, err = os.Stat("go.mod")
		if err != nil {
			lib.RapiError("go.mod not found, is this an existing project?")
			lib.RapiExitBad()
		}

		file, err := os.ReadFile("go.mod")
		lib.RapiErrorCheck(err)
		found := false
		foundFramework := ""
		for _, framework := range supportedFrameworks {
			if strings.Contains(string(file), framework) {
				found = true
				foundFramework = framework
				break
			}
		}
		if !found {
			lib.RapiError("No supported framework found in go.mod, supported frameworks are:")
			for _, framework := range supportedFrameworks {
				println(framework)
			}
			lib.RapiExitBad()
		}
		projectName := ""
		nameActive := false
		for _, char := range string(file) {
			if string(char) == " " {
				nameActive = true
			}
			if nameActive {
				projectName += string(char)
			}
			//dosnt work
			if char == 92 { // / in ascii
				nameActive = false
				break
			}
		}
		if projectName == "" {
			lib.RapiError("No module name found in go.mod")
			lib.RapiExitBad()
		}

		var routesPath string
		lib.RapiInfo("Where are your routes located?")
		for routesPath == "" {
			_, err = fmt.Scanln(&routesPath)
			lib.RapiErrorCheck(err)
		}
		_, err = os.Stat(routesPath)
		if err != nil {
			lib.RapiError("Invalid path to routes")
			lib.RapiExitBad()
		}
		var middlewarePath string
		lib.RapiInfo("Where are your middlewares located?")
		for middlewarePath == "" {
			_, err = fmt.Scanln(&middlewarePath)
			lib.RapiErrorCheck(err)
		}
		_, err = os.Stat(middlewarePath)
		if err != nil {
			lib.RapiError("Invalid path to middlewares")
			lib.RapiExitBad()
		}

		viper.AddConfigPath(".")
		viper.SetConfigName("rapi")
		viper.SetConfigType("json")
		viper.Set("projectName", projectName)
		viper.Set("framework", foundFramework)
		viper.Set("routesPath", routesPath)
		viper.Set("middlewarePath", middlewarePath)
		err = viper.SafeWriteConfig()
		lib.RapiErrorCheck(err)

		lib.RapiInfo("RAPI initialized successfully")
	},
}
