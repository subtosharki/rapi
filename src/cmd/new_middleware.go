package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/subtosharki/rapi/src/lib"
	"github.com/subtosharki/rapi/src/templates/fiber"
	"github.com/subtosharki/rapi/src/templates/gin"
	"os"
	"strings"
	"unicode"
)

func init() {
	Root.AddCommand(newMiddlewareCmd)
}

var newMiddlewareCmd = &cobra.Command{
	Use:   "new:middleware",
	Short: "Create a new middleware",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat("rapi.json")
		if err != nil {
			lib.Error("rapi.json not found, please run rapi init")
			lib.ExitBad()
		}
		lib.LoadConfig()
		middlewaresPath := viper.GetString("middlewarespath")
		if middlewaresPath == "" {
			lib.Error("middlewarespath not found in rapi.json")
		}
		framework := viper.GetString("framework")
		if framework == "" {
			lib.Error("framework not found in rapi.json")
		}
		middlewareName := args[0]

		_, err = os.Stat(middlewaresPath)
		if err != nil {
			lib.Error("middlewarespath not found")
			lib.ExitBad()
		}

		if strings.Contains(middlewareName, "/") {
			lib.Error("Middleware name cannot contain /")
			lib.ExitBad()
		}
		_, err = os.Stat(middlewaresPath + "/" + middlewareName + ".go")
		if err == nil {
			lib.Error("Middleware already exists")
			lib.ExitBad()
		}
		file, err := os.Create(middlewaresPath + "/" + middlewareName + ".go")
		if err != nil {
			lib.Error("Error creating route file")
			lib.ExitBad()
		}

		pathName := strings.Split(middlewaresPath, "/")

		runes := []rune(middlewareName)
		runes[0] = unicode.ToUpper(runes[0])
		middlewareName = string(runes)
		switch framework {
		case "fiber":
			_, err := file.WriteString(fiber.BasicRoute(middlewareName, pathName[len(pathName)-1]))
			if err != nil {
				lib.Error("Error writing to file")
				lib.ExitBad()
			}
		case "gin":
			_, err := file.WriteString(gin.BasicRoute(middlewareName, pathName[len(pathName)-1]))
			if err != nil {
				lib.Error("Error writing to file")
				lib.ExitBad()
			}
		default:
			lib.Error("Invalid framework")
			lib.ExitBad()
		}
		err = file.Close()
		if err != nil {
			lib.Error("Error closing file")
			lib.ExitBad()
		}
		lib.Info("Route created successfully, make sure to add it to your main file")
		lib.ExitOk()
	},
}
