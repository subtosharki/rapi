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
	Short: "Initialize rapi in a current project",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat("rapi.json")
		if err == nil {
			lib.Error("rapi.json found, already initialized")
			lib.ExitBad()
		}
		_, err = os.Stat("go.mod")
		if err != nil {
			lib.Error("go.mod not found, is this an existing project?")
			lib.ExitBad()
		}

		file, err := os.ReadFile("go.mod")
		lib.ErrorCheck(err)
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
			lib.Error("No supported framework found in go.mod, supported frameworks are:")
			for _, framework := range supportedFrameworks {
				println(framework)
			}
			lib.ExitBad()
		}
		if strings.Split(string(file), " ")[0] != "module" {
			lib.Error("Invalid go.mod file")
			lib.ExitBad()
		}
		projectName := strings.Split(strings.Split(string(file), " ")[1], "\n")[0]
		var routesPath string
		commonPaths := []string{"src/routes", "routes", "src/route", "route"}
		for _, path := range commonPaths {
			_, err = os.Stat(path)
			if err == nil {
				lib.Info("Using routes in " + path)
				routesPath = path
				break
			}
		}
		if routesPath == "" {
			lib.Info("Where are your routes located?")
			for routesPath == "" {
				_, err = fmt.Scanln(&routesPath)
				lib.ErrorCheck(err)
			}
			_, err = os.Stat(routesPath)
			if err != nil {
				lib.Error("Invalid path to routes")
				lib.ExitBad()
			}
		}
		var middlewaresPath string
		commonPaths = []string{"src/middlewares", "middlewares", "src/middleware", "middleware"}
		for _, path := range commonPaths {
			_, err = os.Stat(path)
			if err == nil {
				lib.Info("Using middlewares in " + path)
				middlewaresPath = path
				break
			}
		}
		if middlewaresPath == "" {
			lib.Info("Where are your middlewares located?")
			for middlewaresPath == "" {
				_, err = fmt.Scanln(&middlewaresPath)
				lib.ErrorCheck(err)
			}
			_, err = os.Stat(middlewaresPath)
			if err != nil {
				lib.Error("Invalid path to middlewares")
				lib.ExitBad()
			}
		}
		viper.AddConfigPath(".")
		viper.SetConfigName("rapi")
		viper.SetConfigType("json")
		viper.Set("projectname", projectName)
		viper.Set("framework", foundFramework)
		viper.Set("routespath", routesPath)
		viper.Set("middlewarespath", middlewaresPath)
		err = viper.SafeWriteConfig()
		lib.ErrorCheck(err)

		lib.Info("Initialized successfully")
		lib.ExitOk()
	},
}
