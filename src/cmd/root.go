package cmd

import (
	"github.com/spf13/cobra"
	"github.com/subtosharki/rapi/src/lib"
)

var Root = &cobra.Command{
	Use: "rapi",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		lib.ErrorCheck(err)
	},
}
