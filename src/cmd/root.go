// certificateManager
// src/cmd/root.go

package cmd

import (
	"fmt"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"github.com/spf13/cobra"
	"os"
	"runtime"
	"vclt/env"
	"vclt/kv"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "vclt",
	Short:   "Hashicorp Vault client",
	Version: hf.White(fmt.Sprintf("1.00.00-0-%s (2024.06.27)", runtime.GOARCH)),
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
		if _, err := kv.Login(); err != nil {
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
	rootCmd.AddCommand(envCmd, kvCmd, loginCmd)

	envCmd.AddCommand(envListCmd, envRmCmd, envAddCmd, envInfoCmd)

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
1.00.00		2024.06.27		Initial version
\n`)
}
