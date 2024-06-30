// vclt
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : sysCommands.go
// Original timestamp : 2024/06/30 18:44

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var sysCmd = &cobra.Command{
	Use:     "kv",
	Example: "vclt kv { get | put | add | list }",
	Short:   "kv store sub-command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You need to specify one of the following subcommand: add | delete | verify | list")
		os.Exit(0)
	},
}
