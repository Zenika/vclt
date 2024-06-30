// vclt
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/cmd/kvCommands.go
// Original timestamp: 2024/06/28 14:20

package cmd

import (
	"fmt"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"github.com/spf13/cobra"
	"os"
	"strconv"
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
		nVer := 0
		if len(args) < 2 {
			fmt.Println("Usage: vclt PATH FIELD VERSION")
		}
		if len(args) > 2 {
			nVer, _ = strconv.Atoi(args[2])
		}
		if res, ce := kv.Get(args[0], args[1], nVer); ce != nil {
			ce.Error()
		} else {
			if !kv.Quiet {
				fmt.Printf("%s: %s\n", args[1], hf.Green(fmt.Sprintf("%s", res)))
			} else {
				fmt.Printf("%s\n", res)
			}
		}

		os.Exit(0)
	},
}

func init() {
	kvCmd.AddCommand(kvGetCmd)

	kvCmd.PersistentFlags().BoolVarP(&kv.Quiet, "quiet", "q", false, "Only displays the value from the fetched value(s), defaults to FALSE")
}
