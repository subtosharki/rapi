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
	Root.AddCommand(newRouteCmd)
}

var newRouteCmd = &cobra.Command{
	Use:   "new:route",
	Short: "Create a new route",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat("rapi.json")
		if err != nil {
			lib.Error("rapi.json not found, please run rapi init")
			lib.ExitBad()
		}
		lib.LoadConfig()
		routesPath := viper.GetString("routespath")
		if routesPath == "" {
			lib.Error("routespath not found in rapi.json")
		}
		framework := viper.GetString("framework")
		if framework == "" {
			lib.Error("framework not found in rapi.json")
		}
		routeName := args[0]

		_, err = os.Stat(routesPath)
		if err != nil {
			lib.Error("routespath not found")
			lib.ExitBad()
		}

		if strings.Contains(routeName, "/") {
			lib.Error("Route name cannot contain /")
			lib.ExitBad()
		}
		_, err = os.Stat(routesPath + "/" + routeName + ".go")
		if err == nil {
			lib.Error("Route already exists")
			lib.ExitBad()
		}
		file, err := os.Create(routesPath + "/" + routeName + ".go")
		if err != nil {
			lib.Error("Error creating route file")
			lib.ExitBad()
		}

		pathName := strings.Split(routesPath, "/")

		runes := []rune(routeName)
		runes[0] = unicode.ToUpper(runes[0])
		routeName = string(runes)
		switch framework {
		case "fiber":
			_, err := file.WriteString(fiber.BasicRoute(routeName, pathName[len(pathName)-1]))
			if err != nil {
				lib.Error("Error writing to file")
				lib.ExitBad()
			}
		case "gin":
			_, err := file.WriteString(gin.BasicRoute(routeName, pathName[len(pathName)-1]))
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
		lib.Info("Route created successfully, make sure to add it to your routes.go file")
		lib.ExitOk()
	},
}
