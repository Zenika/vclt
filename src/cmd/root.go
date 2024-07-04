// certificateManager
// src/cmd/root.go

package cmd

import (
	"fmt"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"github.com/spf13/cobra"
	"os"
	"runtime"
	"time"
	"vclt/env"
	"vclt/sys"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "vclt",
	Short:   "Hashicorp Vault client",
	Version: hf.White(fmt.Sprintf("1.01.00-1-%s (%s)", runtime.GOARCH, time.Now().Format("2006.01.02"))),
}

var clCmd = &cobra.Command{
	Use:     "changelog",
	Aliases: []string{"cl"},
	Short:   "Shows changelog",
	Run: func(cmd *cobra.Command, args []string) {
		changelog()
	},
}

var loginCmd = &cobra.Command{
	Use:     "login",
	Aliases: []string{"auth"},
	Short:   "Login to Vault, using the userpass method",
	Run: func(cmd *cobra.Command, args []string) {
		if _, _, err := sys.Login(true); err != nil {
			err.Error()
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.DisableAutoGenTag = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(clCmd)
	rootCmd.AddCommand(envCmd, kvCmd, loginCmd, sysCmd)

	rootCmd.PersistentFlags().StringVarP(&env.ConfigFile, "env", "e", "defaultEnv.json", "Default env configuration file; this is a per-user setting.")
}

func changelog() {
	//fmt.Printf("\x1b[2J")
	fmt.Printf("\x1bc")

	fmt.Println("CHANGELOG")
	fmt.Println()
	fmt.Println()

	fmt.Print(`
VERSION		DATE			COMMENT
-------		----			-------
1.01.00		2024.07.02		Completed kv get, kv lse, kv lsf
1.00.00		2024.06.27		Initial version
\n`)
}
