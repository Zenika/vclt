// vclt
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/cmd/kv.go
// Original timestamp: 2024/06/28 14:20

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"vclt/kv"
)

var kvCmd = &cobra.Command{
	Use:     "kv",
	Example: "vclt kv { get | put | add | list }",
	Short:   "kv store sub-command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You need to specify one of the following subcommand: add | delete | verify | list")
		os.Exit(0)
	},
}

var kvGetCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"read"},
	//Example: "vclt kv { get | put | add | list }",
	Short: "Read an entry in a secret",
	Run: func(cmd *cobra.Command, args []string) {
		kv.Get("")
		os.Exit(0)
	},
}

func init() {
	kvCmd.AddCommand(kvGetCmd)
}
