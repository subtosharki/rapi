package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/subtosharki/rapi/src/lib"
	"os"
)

func init() {
	Root.AddCommand(newRouteCmd)
}

var newRouteCmd = &cobra.Command{
	Use:   "new:route",
	Short: "Create a new route",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat("rapi.json")
		if err != nil {
			lib.RapiError("rapi.json not found, please run rapi init")
			lib.RapiExitBad()
		}
		routesPath := viper.GetString("routesPath")
		if routesPath == "" {
			lib.RapiError("routesPath not found in rapi.json")
		}
		framework := viper.GetString("framework")
		if framework == "" {
			lib.RapiError("framework not found in rapi.json")
		}
		//routeName := args[0]

	},
}
