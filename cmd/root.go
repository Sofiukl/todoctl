/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/sofiukl/todoctl/utils"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todoctl",
	Short: `todoctl is a CLI client of task control.`,
	Long: `todoctl is a CLI client of task control. you can use our task control portal
	 also for managing your tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set command values based on preference
		utils.SetCommandValues(cmd)
		serverUrl, _ := cmd.Flags().GetString("server-url")
		fmt.Println("server url is:", serverUrl)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Add flags in root command if required
	rootCmd.Flags().StringP("server-url", "u", "http://localhost:5000", "Should come from flag first, then the config file, then the default last")

	// Initialize config variables with viper
	viper.BindPFlag("server_url", rootCmd.Flags().Lookup("server-url"))
	cobra.OnInitialize(utils.InitConfig)
}
