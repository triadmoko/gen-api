/*
Copyright © 2023 NAME HERE triadmoko12@gmail.com
*/
package create

import (
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var Create = &cobra.Command{
	Use:   "create",
	Short: "command create your project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			cmd.PrintErr(err)
		}
	},
}
